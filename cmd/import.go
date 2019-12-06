/*
Copyright © 2019 Michael Zion noisleahcim@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Short Import CMD Description",
	Long:  `Long Import CMD Description`,
	Run: func(cmd *cobra.Command, args []string) {
		fileUrl := "https://raw.githubusercontent.com/noisleahcim/shlib/master/lib/logging.sh"
		dirPath := ".shlib~"
		filePath := "logging.sh"
		fullFilePath := dirPath + "/" + filePath

		if err := createTempModulesDir(dirPath); err != nil {
			panic(err)
		}

		if err := downloadFile(fullFilePath, fileUrl); err != nil {
			panic(err)
		}

		if err := printModule(fullFilePath); err != nil {
			panic(err)
		}

		if err := deleteTempModulesDir(dirPath); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolP("all", "a", true, "Help message for --all")
}

func createTempModulesDir(dirPath string) error {
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModeDir)
	}

	return err
}

func deleteTempModulesDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	return err
}

func downloadFile(filePath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func printModule(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	fmt.Print(content)
	return err
}
