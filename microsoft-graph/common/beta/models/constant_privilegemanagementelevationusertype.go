package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationUserType string

const (
	PrivilegeManagementElevationUserTypeazureAd      PrivilegeManagementElevationUserType = "AzureAd"
	PrivilegeManagementElevationUserTypehybrid       PrivilegeManagementElevationUserType = "Hybrid"
	PrivilegeManagementElevationUserTypelocal        PrivilegeManagementElevationUserType = "Local"
	PrivilegeManagementElevationUserTypeundetermined PrivilegeManagementElevationUserType = "Undetermined"
)

func PossibleValuesForPrivilegeManagementElevationUserType() []string {
	return []string{
		string(PrivilegeManagementElevationUserTypeazureAd),
		string(PrivilegeManagementElevationUserTypehybrid),
		string(PrivilegeManagementElevationUserTypelocal),
		string(PrivilegeManagementElevationUserTypeundetermined),
	}
}

func parsePrivilegeManagementElevationUserType(input string) (*PrivilegeManagementElevationUserType, error) {
	vals := map[string]PrivilegeManagementElevationUserType{
		"azuread":      PrivilegeManagementElevationUserTypeazureAd,
		"hybrid":       PrivilegeManagementElevationUserTypehybrid,
		"local":        PrivilegeManagementElevationUserTypelocal,
		"undetermined": PrivilegeManagementElevationUserTypeundetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationUserType(input)
	return &out, nil
}
