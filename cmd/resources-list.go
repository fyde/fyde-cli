/*
Copyright © 2019 Fyde, Inc. <hello@fyde.com>

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
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	apiresources "github.com/oNaiPs/fyde-cli/client/access_resources"
	"github.com/oNaiPs/fyde-cli/models"
)

// resourcesListCmd represents the list command
var resourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := preRunCheckAuth(cmd, args)
		if err != nil {
			return err
		}

		err = preRunFlagCheckOutput(cmd, args)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		uparams := apiresources.NewListResourcesParams()
		completePayload := []*apiresources.ListResourcesOKBodyItems0{}
		err := forAllPages(uparams, func() (int64, error) {
			resp, err := global.Client.AccessResources.ListResources(uparams, global.AuthWriter)
			if err == nil {
				completePayload = append(completePayload, resp.Payload...)
			}
			return resp.Total, err
		})
		if err != nil {
			return processErrorResponse(err)
		}

		tw := table.NewWriter()
		tw.Style().Format.Header = text.FormatDefault
		tw.AppendHeader(table.Row{
			"ID",
			"Name",
			"Public host",
			"Access policy",
			"Port ext:int",
			"Access Proxy",
		})

		for _, item := range completePayload {
			accessPolicies := strings.Join(funk.Map(item.AccessPolicies, func(g *models.AccessResourceAccessPoliciesItems0) string {
				return g.Name
			}).([]string), ",")

			tw.AppendRow(table.Row{
				item.ID,
				item.Name,
				item.PublicHost,
				accessPolicies,
				strings.Join(item.Ports, ","),
				item.AccessProxyID,
			})
		}

		var result string
		outputFormat, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		switch outputFormat {
		case "table":
			result = tw.Render()
		case "csv":
			result = tw.RenderCSV()
		case "json":
			result = renderJSON(completePayload)
		}
		fmt.Println(result)
		return nil
	},
}

func init() {
	resourcesCmd.AddCommand(resourcesListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourcesListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourcesListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	resourcesListCmd.Flags().StringP("filter", "f", "", "filter resources")
	resourcesListCmd.Flags().StringP("output", "o", "table", "output format (table, json or csv)")
}
