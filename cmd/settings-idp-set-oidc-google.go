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
	api "github.com/barracuda-cloudgen-access/access-cli/client/identity_providers"
	"github.com/spf13/cobra"
)

// setGoogleIdpCmd represents the get command
var setGoogleIdpCmd = &cobra.Command{
	Use:   "google",
	Short: "Set google idp configuration",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := preRunCheckAuth(cmd, args)
		if err != nil {
			return err
		}

		err = preRunFlagChecks(cmd, args)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		params := api.NewCreateIdentityProviderParams()
		setTenant(cmd, params)

		cmd.SilenceUsage = true // errors beyond this point are no longer due to malformed input

		body := api.CreateIdentityProviderBody{
			IdentityProvider: &api.CreateIdentityProviderParamsBodyIdentityProvider{
				IdpType: "google_oidc",
			}}
		params.SetIdentityProvider(body)

		resp, err := global.Client.IdentityProviders.CreateIdentityProvider(params, global.AuthWriter)
		if err != nil {
			return err
		}

		tw := identityProviderConfigBuildTableWriter()
		if resp.Payload.ID > 0 {
			identityProviderTableWriterAppend(tw, *resp.Payload)
		}
		return printListOutputAndError(cmd, resp.Payload, tw, 1, err)
	},
}

func init() {
	setIdpCmd.AddCommand(setGoogleIdpCmd)

	initOutputFlags(setGoogleIdpCmd)
	initTenantFlags(setGoogleIdpCmd)
}
