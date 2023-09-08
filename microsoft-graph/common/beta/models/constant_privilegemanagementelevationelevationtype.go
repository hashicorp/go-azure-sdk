package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationElevationType string

const (
	PrivilegeManagementElevationElevationTypesupportApprovedElevation PrivilegeManagementElevationElevationType = "SupportApprovedElevation"
	PrivilegeManagementElevationElevationTypeundetermined             PrivilegeManagementElevationElevationType = "Undetermined"
	PrivilegeManagementElevationElevationTypeunmanagedElevation       PrivilegeManagementElevationElevationType = "UnmanagedElevation"
	PrivilegeManagementElevationElevationTypeuserConfirmedElevation   PrivilegeManagementElevationElevationType = "UserConfirmedElevation"
	PrivilegeManagementElevationElevationTypezeroTouchElevation       PrivilegeManagementElevationElevationType = "ZeroTouchElevation"
)

func PossibleValuesForPrivilegeManagementElevationElevationType() []string {
	return []string{
		string(PrivilegeManagementElevationElevationTypesupportApprovedElevation),
		string(PrivilegeManagementElevationElevationTypeundetermined),
		string(PrivilegeManagementElevationElevationTypeunmanagedElevation),
		string(PrivilegeManagementElevationElevationTypeuserConfirmedElevation),
		string(PrivilegeManagementElevationElevationTypezeroTouchElevation),
	}
}

func parsePrivilegeManagementElevationElevationType(input string) (*PrivilegeManagementElevationElevationType, error) {
	vals := map[string]PrivilegeManagementElevationElevationType{
		"supportapprovedelevation": PrivilegeManagementElevationElevationTypesupportApprovedElevation,
		"undetermined":             PrivilegeManagementElevationElevationTypeundetermined,
		"unmanagedelevation":       PrivilegeManagementElevationElevationTypeunmanagedElevation,
		"userconfirmedelevation":   PrivilegeManagementElevationElevationTypeuserConfirmedElevation,
		"zerotouchelevation":       PrivilegeManagementElevationElevationTypezeroTouchElevation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationElevationType(input)
	return &out, nil
}
