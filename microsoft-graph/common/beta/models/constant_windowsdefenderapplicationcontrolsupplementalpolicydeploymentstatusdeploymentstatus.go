package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus string

const (
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusnotAuthorizedByToken WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "NotAuthorizedByToken"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatuspolicyNotFound       WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "PolicyNotFound"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatussuccess              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "Success"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatustokenError           WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "TokenError"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusunknown              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "Unknown"
)

func PossibleValuesForWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus() []string {
	return []string{
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusnotAuthorizedByToken),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatuspolicyNotFound),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatussuccess),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatustokenError),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusunknown),
	}
}

func parseWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus(input string) (*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus, error) {
	vals := map[string]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus{
		"notauthorizedbytoken": WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusnotAuthorizedByToken,
		"policynotfound":       WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatuspolicyNotFound,
		"success":              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatussuccess,
		"tokenerror":           WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatustokenError,
		"unknown":              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus(input)
	return &out, nil
}
