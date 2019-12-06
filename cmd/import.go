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
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Short Import CMD Description",
	Long:  `Long Import CMD Description`,
	Run: func(cmd *cobra.Command, args []string) {
		importMain()
	},
}

func importMain() {
	// variables
	fileUrl := "https://raw.githubusercontent.com/noisleahcim/shlib/master/lib/logging.sh"
	dirPath := ".shlib"
	filePath := "logging.sh"
	fullFilePath := "./" + dirPath + "/" + filePath

	// main
	createTempModulesDir(dirPath)
	downloadFile(fullFilePath, fileUrl)
	sourceCode(fullFilePath)
	deleteTempModulesDir(dirPath)
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolP("all", "a", true, "Help message for --all")
}

func createTempModulesDir(dirPath string) {
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}
}

func deleteTempModulesDir(dirPath string) {
	err := os.RemoveAll(dirPath)

	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(filePath string, url string) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	os.Chmod(filePath, 0744)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}

func sourceCode(filePath string) {
	cmd := exec.Command("/bin/sh", "-c", filePath, "&&", "log_info 'bla'")

	var out bytes.Buffer
	cmd.Stdout = &out

	fmt.Println(cmd.Stdout)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
