package tenantaccessgit

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessIdName string

const (
	AccessIdNameAccess    AccessIdName = "access"
	AccessIdNameGitAccess AccessIdName = "gitAccess"
)

func PossibleValuesForAccessIdName() []string {
	return []string{
		string(AccessIdNameAccess),
		string(AccessIdNameGitAccess),
	}
}
