package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate string

const (
	OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreateguest  OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate = "Guest"
	OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreatemember OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate = "Member"
)

func PossibleValuesForOnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate() []string {
	return []string{
		string(OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreateguest),
		string(OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreatemember),
	}
}

func parseOnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate(input string) (*OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate, error) {
	vals := map[string]OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate{
		"guest":  OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreateguest,
		"member": OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreatemember,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate(input)
	return &out, nil
}
