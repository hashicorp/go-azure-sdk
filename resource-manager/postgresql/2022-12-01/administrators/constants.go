package administrators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrincipalType string

const (
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUnknown          PrincipalType = "Unknown"
	PrincipalTypeUser             PrincipalType = "User"
)

func PossibleValuesForPrincipalType() []string {
	return []string{
		string(PrincipalTypeGroup),
		string(PrincipalTypeServicePrincipal),
		string(PrincipalTypeUnknown),
		string(PrincipalTypeUser),
	}
}
