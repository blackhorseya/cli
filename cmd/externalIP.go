/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// externalIPCmd represents the externalIP command
var externalIPCmd = &cobra.Command{
	Use:   "externalIP",
	Short: "Looks up the External IP addresses",
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest(http.MethodGet, "https://ifconfig.co", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		c := &http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(strings.Trim(string(body), "\n"))

		if copied, err := cmd.Flags().GetBool("copy"); err != nil {
			fmt.Println(err)
			return
		} else {
			if copied {
				if err := clipboard.WriteAll(strings.Trim(string(body), "\n")); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(externalIPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// externalIPCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// externalIPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	externalIPCmd.Flags().Bool("copy", false, "write output to clipboard")
}
