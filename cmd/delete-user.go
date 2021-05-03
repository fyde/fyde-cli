// Package cmd implements access-cli commands
package cmd

/*
Copyright © 2020 Barracuda Networks, Inc.

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

import (
	"fmt"

	"github.com/spf13/cobra"

	apiusers "github.com/barracuda-cloudgen-access/access-cli/client/users"
)

// userDeleteCmd represents the delete command
var userDeleteCmd = &cobra.Command{
	Use:     "delete [user ID]...",
	Aliases: []string{"remove", "rm"},
	Short:   "Delete users",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := preRunCheckAuth(cmd, args)
		if err != nil {
			return err
		}

		err = preRunFlagChecks(cmd, args)
		if err != nil {
			return err
		}

		if !multiOpCheckArgsPresent(cmd, args) {
			return fmt.Errorf("missing user ID argument")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		userIDs, err := multiOpParseInt64Args(cmd, args, "id")
		if err != nil {
			return err
		}

		delete := func(ids []int64) error {
			params := apiusers.NewDeleteUserParams()
			params.SetID(ids)

			_, err = global.Client.Users.DeleteUser(params, global.AuthWriter)
			if err != nil {
				return processErrorResponse(err)
			}
			return nil
		}

		tw, j := multiOpBuildTableWriter()

		if loopControlContinueOnError(cmd) {
			// then we must delete individually, because on a request for multiple deletions,
			// the server does nothing if one fails

			for _, id := range userIDs {
				err = delete([]int64{id})
				var result interface{}
				result = "success"
				if err != nil {
					result = err
				}
				multiOpTableWriterAppend(tw, &j, id, result)
			}
			err = nil
		} else {
			err = delete(userIDs)
			var result interface{}
			result = "success"
			if err != nil {
				result = err
			}
			multiOpTableWriterAppend(tw, &j, "*", result)
		}

		return printListOutputAndError(cmd, j, tw, len(userIDs), err)
	},
}

func init() {
	usersCmd.AddCommand(userDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initMultiOpArgFlags(userDeleteCmd, "user", "delete", "id", "[]int64")
	initOutputFlags(userDeleteCmd)
	initLoopControlFlags(userDeleteCmd)
}
