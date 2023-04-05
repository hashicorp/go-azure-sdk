package accesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyRole string

const (
	AccessPolicyRoleContributor AccessPolicyRole = "Contributor"
	AccessPolicyRoleReader      AccessPolicyRole = "Reader"
)

func PossibleValuesForAccessPolicyRole() []string {
	return []string{
		string(AccessPolicyRoleContributor),
		string(AccessPolicyRoleReader),
	}
}
