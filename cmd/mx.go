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
	"net"

	"github.com/spf13/cobra"
)

// mxCmd represents the mx command
var mxCmd = &cobra.Command{
	Use:   "mx",
	Short: "Looks up the MX for a particular host",
	RunE: func(cmd *cobra.Command, args []string) error {
		host := cmd.Flag("host").Value.String()
		if mx, err := net.LookupMX(host); err != nil {
			return err
		} else {
			for _, m := range mx {
				fmt.Println(m.Host, m.Pref)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(mxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mxCmd.Flags().String("host", "seancheng.space", "Particular host")
}
