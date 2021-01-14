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

// cnameCmd represents the cname command
var cnameCmd = &cobra.Command{
	Use:   "cname",
	Short: "Looks up the CNAME for a particular host",
	RunE: func(cmd *cobra.Command, args []string) error {
		host := cmd.Flag("host").Value.String()
		if cname, err := net.LookupCNAME(host); err != nil {
			return err
		} else {
			fmt.Println(cname)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cnameCmd.Flags().String("host", "seancheng.space", "Particular host")
}
