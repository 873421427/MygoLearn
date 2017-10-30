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
	"strings"

	"Agenda/entity"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var createMeetingsCmd = &cobra.Command{
	Use:   "createMeetings",
	Short: "Create meetings.",
	Long: `If you want to create a meeting, you should declare the title name,
		which can't be the same as others' title name, some participators(at least one),
		start time as the format of (yyyy-mm-dd), and end time as the format of (yyyy-mm-dd).`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("title", title)

		participatorStr, _ := comd.Flags().GetString("participators")
		checkEmpty("participators", participatorStr)
		participators := strings.Split(participatorStr, " ")

		startTime, _ := comd.Flags().GetString("start")
		checkEmpty("Start Time", startTime)

		endTime, _ := comd.Flags().GetString("end")
		checkEmpty("End Time", endTime)

		if err := entity.MeetingCreate(title, participators, startTime, endTime); err != nil {
			println(err)
		} else {
			println("A meeting was successfully created!")
		}
	},
}

var addParticipatorCmd = &cobra.Command{
	Use:   "addParticipator",
	Short: "Add your own meetings' participators.",
	Long: `You can append some participators from your own meeting
	by specifying the title name.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("title", title)

		participatorStr, _ := comd.Flags().GetString("participators")
		checkEmpty("participators", participatorStr)

		participators := strings.Split(participatorStr, ",")

		if err := entity.AddMeetingParticipator(title, participators); err != nil {
			println(err)
		} else {
			println("Add meeting participator successfully!")
		}
	},
}

var rmParticipatorCmd = &cobra.Command{
	Use:   "rmParticipator",
	Short: "Remove your own meetings' participators.",
	Long: `You can remove some participators from your own meeting
	by specifying the title name.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("title", title)

		participatorStr, _ := comd.Flags().GetString("participators")
		checkEmpty("participators", participatorStr)

		participators := strings.Split(participatorStr, ",")

		if err := entity.RemoveParticipator(title, participators); err != nil {
			println(err)
		} else {
			println("Remove meeting participator successfully!")
		}
	},
}

var listMeetingsCmd = &cobra.Command{
	Use:   "listMeetings",
	Short: "List all of your own meetings during a time interval.",
	Long: `You can see the detail information of all of meetings,
	which you attended, during a time interval.`,
	Run: func(comd *cobra.Command, args []string) {
		startTime, _ := comd.Flags().GetString("start")
		checkEmpty("Start Time", startTime)

		endTime, _ := comd.Flags().GetString("end")
		checkEmpty("End Time", endTime)

		if err := entity.ListMeeting(startTime, endTime); err != nil {
			println(err)
		} else {
			println("Listing meeting operation completed successfully!")
		}
	},
}

var delAMeetingCmd = &cobra.Command{
	Use:   "deleteAMeeting",
	Short: "Cancel your own meeting by specifying title name.",
	Long:  `Using this command, you are able to cancel the meetings, which are created by you.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("Title", title)

		if err := entity.DeleteAMeeting(title); err != nil {
			println(err)
		} else {
			println("The meeting was successfully deleted!")
		}
	},
}

var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "Quit meetings.",
	Long:  `You can quit any meetings you want, which are you attended, not created.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("Title", title)

		if err := entity.QuitMeeting(title); err != nil {
			println(err)
		} else {
			println("You've successfully quit the meeting " + title + "!")
		}
	},
}

/*delete all meetings which are sponsored by this login user*/
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all meetings you attended or created.",
	Long:  `Using this command, you can clear all of the meetings you attended or created.`,
	Run: func(comd *cobra.Command, args []string) {

		if err := entity.DeleteAllMeeting(); err != nil {
			println(err)
		} else {
			println("You've successfully cleared all the meetings you sponsored!")
		}
	},
}

func init() {
	RootCmd.AddCommand(createMeetingsCmd)
	createMeetingsCmd.Flags().StringP("title", "t", "", "Input title name.")
	createMeetingsCmd.Flags().StringP("participators", "p", "", "Input participator name.")
	createMeetingsCmd.Flags().StringP("start", "s", "", "Input start time as the format of (yyyy-mm-dd).")
	createMeetingsCmd.Flags().StringP("end", "e", "", "Input end time as the format of (yyyy-mm-dd).")

	RootCmd.AddCommand(addParticipatorCmd)
	addParticipatorCmd.Flags().StringP("title", "t", "", "Input the title name.")
	addParticipatorCmd.Flags().StringP("participators", "p", "", "Input the participators.")

	RootCmd.AddCommand(rmParticipatorCmd)
	rmParticipatorCmd.Flags().StringP("title", "t", "", "Input the title name.")
	rmParticipatorCmd.Flags().StringP("participators", "p", "", "Input the participators.")

	RootCmd.AddCommand(listMeetingsCmd)
	listMeetingsCmd.Flags().StringP("start", "s", "", "Input the start time.(yyyy-mm-dd)")
	listMeetingsCmd.Flags().StringP("end", "e", "", "Input the end time.(yyyy-mm-dd)")

	RootCmd.AddCommand(delAMeetingCmd)
	delAMeetingCmd.Flags().StringP("title", "t", "", "Input the title.")

	RootCmd.AddCommand(quitCmd)
	quitCmd.Flags().StringP("title", "t", "", "Input the title.")

	RootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
