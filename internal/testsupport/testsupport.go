/*
 * Copyright 2021 Skyscanner Limited.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * https://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testsupport

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Pwd() string {
	dir, _ := os.Getwd()
	return filepath.Base(dir)
}

func CreateAndEnterTempDirectory() string {
	tempDir, _ := ioutil.TempDir("", "turbolift-test-*")
	err := os.Chdir(tempDir)
	if err != nil {
		panic(err)
	}
	return tempDir
}

func PrepareTempCampaign(createDirs bool, repos ...string) string {
	tempDir := CreateAndEnterTempDirectory()

	delimitedList := strings.Join(repos, "\n")
	err := ioutil.WriteFile("repos.txt", []byte(delimitedList), os.ModePerm|0o644)
	if err != nil {
		panic(err)
	}

	if createDirs {
		for _, name := range repos {
			dirToCreate := path.Join("work", name)
			err := os.MkdirAll(dirToCreate, os.ModeDir|0o755)
			if err != nil {
				panic(err)
			}
		}
	}

	dummyPrDescription := "# PR title\nPR body"
	err = ioutil.WriteFile("README.md", []byte(dummyPrDescription), os.ModePerm|0o644)
	if err != nil {
		panic(err)
	}

	return tempDir
}

func CreateAnotherRepoFile(filename string, repos ...string) {
	delimitedList := strings.Join(repos, "\n")
	err := ioutil.WriteFile(filename, []byte(delimitedList), os.ModePerm|0o644)
	if err != nil {
		panic(err)
	}
}

func CreateOrUpdatePrDescriptionFile(filename string, prTitle string, prBody string) {
	prDescription := fmt.Sprintf("# %s\n%s", prTitle, prBody)
	err := os.WriteFile(filename, []byte(prDescription), os.ModePerm|0o644)
	if err != nil {
		panic(err)
	}
}

func UseDefaultPrDescription(dirName string) {
	originalPrTitle := fmt.Sprintf("TODO: Title of Pull Request (%s)", dirName)
	originalPrBody := `TODO: This file will serve as both a README and the description of the PR. Describe the pull request using markdown in this file. Make it clear why the change is being made, and make suggestions for anything that the reviewer may need to do.

By approving this PR, you are confirming that you have adequately and effectively reviewed this change.

## How this change was made
TODO: Describe the approach that was used to select repositories for this change
TODO: Describe any shell commands, scripts, manual operations, etc, that were used to make changes

<!-- Please keep the footer below, to support turbolift usage tracking -->
<sub>This PR was generated using [turbolift](https://github.com/Skyscanner/turbolift).</sub>`
	CreateOrUpdatePrDescriptionFile("README.md", originalPrTitle, originalPrBody)
}
