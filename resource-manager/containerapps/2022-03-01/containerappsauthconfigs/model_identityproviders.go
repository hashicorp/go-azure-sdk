package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviders struct {
	Apple                        *Apple                                  `json:"apple"`
	AzureActiveDirectory         *AzureActiveDirectory                   `json:"azureActiveDirectory"`
	AzureStaticWebApps           *AzureStaticWebApps                     `json:"azureStaticWebApps"`
	CustomOpenIdConnectProviders *map[string]CustomOpenIdConnectProvider `json:"customOpenIdConnectProviders,omitempty"`
	Facebook                     *Facebook                               `json:"facebook"`
	GitHub                       *GitHub                                 `json:"gitHub"`
	Google                       *Google                                 `json:"google"`
	Twitter                      *Twitter                                `json:"twitter"`
}
