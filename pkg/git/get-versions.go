package git

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pull-request-formatter/pkg/config"
	"regexp"
	"strconv"
	"strings"
)

func GetVersions() (versions []version, err error) {

	files, err := getFiles()
	if err != nil {
		return
	}

	changedFolders := findFolders(files)

	var reqOld *http.Request
	var reqNew *http.Request
	var resOld *http.Response
	var resNew *http.Response

	for _, folder := range changedFolders {
		oldBranch := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s/%s", config.GitOwner, config.GitRepo, config.FilePathToUpdate, folder, config.FileToUpdate)

		newBranch := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s/%s?ref=%s", config.GitOwner, config.GitRepo, config.FilePathToUpdate, folder, config.FileToUpdate, config.CommitBranch)

		reqOld, err = http.NewRequest("GET", oldBranch, nil)
		if err != nil {
			return
		}

		reqNew, err = http.NewRequest("GET", newBranch, nil)
		if err != nil {
			return
		}

		resOld, err = send(reqOld)
		if err != nil {
			return
		}

		resNew, err = send(reqNew)
		if err != nil {
			return
		}

		var initFileNew fileResponse
		var initFileOld fileResponse

		re := regexp.MustCompile(config.VersionRegex)

		err = json.NewDecoder(resNew.Body).Decode(&initFileNew)
		if err != nil {
			return
		}

		err = resNew.Body.Close()
		if err != nil {
			return
		}

		decodedNew, err := base64.StdEncoding.DecodeString(initFileNew.Content)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resOld.Body).Decode(&initFileOld)
		if err != nil {
			return nil, err
		}

		err = resNew.Body.Close()
		if err != nil {
			return nil, err
		}

		decodedOld, err := base64.StdEncoding.DecodeString(initFileOld.Content)
		if err != nil {
			return nil, err
		}

		matchNew := re.FindStringSubmatch(string(decodedNew))
		matchOld := re.FindStringSubmatch(string(decodedOld))

		if len(matchNew) == 0 || len(matchOld) == 0 {
			fmt.Println("malformed")
		}

		if matchNew[1] != matchOld[1] {
			continue
		}

		oldVersion := matchNew[1]

		newVersion, err := getNextVersion(oldVersion)
		if err != nil {
			return nil, err
		}

		oldIntVersion, err := versionToInt(oldVersion)
		if err != nil {
			return nil, err
		}

		newIntVersion, err := versionToInt(newVersion)
		if err != nil {
			return nil, err
		}

		versions = append(versions, version{
			Name:          folder,
			OldVersion:    oldVersion,
			NewVersion:    newVersion,
			OldIntVersion: oldIntVersion,
			NewIntVersion: newIntVersion,
		})

		newBody := strings.Replace(string(decodedNew), oldVersion, newVersion, 1)

		updateContent(updatedFileBody{
			Sha:     initFileNew.Sha,
			Message: fmt.Sprintf("Version increase for %s from %s to %s", folder, oldVersion, newVersion),
			Committer: committer{
				Name:  config.CommitterName,
				Email: config.CommitterEmail,
			},
			Branch:  config.CommitBranch,
			Content: base64.StdEncoding.EncodeToString([]byte(newBody)),
		}, initFileNew.Path)
	}

	return
}

func findFolders(files []fileBody) (changedFolders []string) {
	pattern := regexp.MustCompile(config.PatternRegex)

	for _, f := range files {
		filepath := f.Filename

		match := pattern.FindString(filepath)

		if match != "" {
			// remove the prefix
			match = strings.TrimPrefix(match, config.FilePathToUpdate+"/")

			if !inArray(changedFolders, match) {
				changedFolders = append(changedFolders, match)
			}
		}
	}
	return
}

func getFiles() (files []fileBody, err error) {
	Init()

	err = getPr()
	if err != nil {
		return
	}

	page := 1

	var file []fileBody

	for true {
		apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%d/files?page=%d", config.GitOwner, config.GitRepo, pr.Number, page)

		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return nil, err
		}

		res, err := send(req)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&file)
		if err != nil {
			return nil, err
		}

		if len(file) == 0 {
			break
		}

		files = append(files, file...)

		page++
	}

	return
}

func getNextVersion(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}

	if patch < 9 {
		patch++
	} else if minor < 9 {
		minor++
		patch = 0
	} else {
		major++
		minor = 0
		patch = 0
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

func versionToInt(version string) (int, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid version format: %s", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("Invalid patch version: %s", parts[2])
	}

	return major*10000 + minor*100 + patch, nil
}

func inArray(array []string, target string) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}
