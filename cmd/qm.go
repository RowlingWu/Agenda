// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// qmCmd represents the qm command
var qmCmd = &cobra.Command{
	Use:   "qm",
	Short: "query meetings from start time to end time",
	Long: `	-sStart_time start time
	-eEnd_time end time
	-a list all meetings from start time to end time
	-m list meetings that you attend from start time to end time`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		my, _ := cmd.Flags().GetBool("my")
		all, _ := cmd.Flags().GetBool("all")
		fmt.Println("start time: " + start)
		fmt.Println("end time: " + end)
		if my {
			fmt.Println("Meetings you attend:")
		}
		if all {
			fmt.Println("All meetings: ")
		}
	},
}

func init() {
	RootCmd.AddCommand(qmCmd)
	qmCmd.Flags().StringP("start", "s", "2017.1.1/00:00", "start time of the meeting")
	qmCmd.Flags().StringP("end", "e", "2017.1.1/01:00", "end time of the meeting")
	qmCmd.Flags().BoolP("my", "m", false, "list meetings that you attend from start time to end time")
	qmCmd.Flags().BoolP("all", "a", false, "list all meetings from start time to end time")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
