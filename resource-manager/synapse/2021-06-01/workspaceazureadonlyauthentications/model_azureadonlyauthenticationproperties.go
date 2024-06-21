package workspaceazureadonlyauthentications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADOnlyAuthenticationProperties struct {
	AzureADOnlyAuthentication bool        `json:"azureADOnlyAuthentication"`
	CreationDate              *string     `json:"creationDate,omitempty"`
	State                     *StateValue `json:"state,omitempty"`
}
