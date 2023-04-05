package vaultcertificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthType string

const (
	AuthTypeAAD                  AuthType = "AAD"
	AuthTypeACS                  AuthType = "ACS"
	AuthTypeAccessControlService AuthType = "AccessControlService"
	AuthTypeAzureActiveDirectory AuthType = "AzureActiveDirectory"
	AuthTypeInvalid              AuthType = "Invalid"
)

func PossibleValuesForAuthType() []string {
	return []string{
		string(AuthTypeAAD),
		string(AuthTypeACS),
		string(AuthTypeAccessControlService),
		string(AuthTypeAzureActiveDirectory),
		string(AuthTypeInvalid),
	}
}
