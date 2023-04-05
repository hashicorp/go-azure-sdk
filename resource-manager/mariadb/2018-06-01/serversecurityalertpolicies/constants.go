package serversecurityalertpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerSecurityAlertPolicyState string

const (
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = "Disabled"
	ServerSecurityAlertPolicyStateEnabled  ServerSecurityAlertPolicyState = "Enabled"
)

func PossibleValuesForServerSecurityAlertPolicyState() []string {
	return []string{
		string(ServerSecurityAlertPolicyStateDisabled),
		string(ServerSecurityAlertPolicyStateEnabled),
	}
}
