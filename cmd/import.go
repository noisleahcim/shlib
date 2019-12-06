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
		fileUrl := "https://raw.githubusercontent.com/noisleahcim/shlib/master/lib/logging.sh"
		filePath := "logging.sh"

		if err := downloadFile(filePath, fileUrl); err != nil {
			panic(err)
		}

		if err := runShellScript(filePath); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolP("all", "a", true, "Help message for --all")
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

func runShellScript(filePath string) error {
	cmd := exec.Command("/bin/sh", filePath, "&&", "source", filePath, "&&", "log_info \"hi\"")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	fmt.Println(cmd.Stdout)
	fmt.Println(cmd.Stderr)

	err := cmd.Run()
	return err
}
