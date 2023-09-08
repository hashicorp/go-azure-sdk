package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequestType string

const (
	AccessPackageAssignmentRequestRequestTypeadminAdd     AccessPackageAssignmentRequestRequestType = "AdminAdd"
	AccessPackageAssignmentRequestRequestTypeadminRemove  AccessPackageAssignmentRequestRequestType = "AdminRemove"
	AccessPackageAssignmentRequestRequestTypeadminUpdate  AccessPackageAssignmentRequestRequestType = "AdminUpdate"
	AccessPackageAssignmentRequestRequestTypenotSpecified AccessPackageAssignmentRequestRequestType = "NotSpecified"
	AccessPackageAssignmentRequestRequestTypeonBehalfAdd  AccessPackageAssignmentRequestRequestType = "OnBehalfAdd"
	AccessPackageAssignmentRequestRequestTypesystemAdd    AccessPackageAssignmentRequestRequestType = "SystemAdd"
	AccessPackageAssignmentRequestRequestTypesystemRemove AccessPackageAssignmentRequestRequestType = "SystemRemove"
	AccessPackageAssignmentRequestRequestTypesystemUpdate AccessPackageAssignmentRequestRequestType = "SystemUpdate"
	AccessPackageAssignmentRequestRequestTypeuserAdd      AccessPackageAssignmentRequestRequestType = "UserAdd"
	AccessPackageAssignmentRequestRequestTypeuserRemove   AccessPackageAssignmentRequestRequestType = "UserRemove"
	AccessPackageAssignmentRequestRequestTypeuserUpdate   AccessPackageAssignmentRequestRequestType = "UserUpdate"
)

func PossibleValuesForAccessPackageAssignmentRequestRequestType() []string {
	return []string{
		string(AccessPackageAssignmentRequestRequestTypeadminAdd),
		string(AccessPackageAssignmentRequestRequestTypeadminRemove),
		string(AccessPackageAssignmentRequestRequestTypeadminUpdate),
		string(AccessPackageAssignmentRequestRequestTypenotSpecified),
		string(AccessPackageAssignmentRequestRequestTypeonBehalfAdd),
		string(AccessPackageAssignmentRequestRequestTypesystemAdd),
		string(AccessPackageAssignmentRequestRequestTypesystemRemove),
		string(AccessPackageAssignmentRequestRequestTypesystemUpdate),
		string(AccessPackageAssignmentRequestRequestTypeuserAdd),
		string(AccessPackageAssignmentRequestRequestTypeuserRemove),
		string(AccessPackageAssignmentRequestRequestTypeuserUpdate),
	}
}

func parseAccessPackageAssignmentRequestRequestType(input string) (*AccessPackageAssignmentRequestRequestType, error) {
	vals := map[string]AccessPackageAssignmentRequestRequestType{
		"adminadd":     AccessPackageAssignmentRequestRequestTypeadminAdd,
		"adminremove":  AccessPackageAssignmentRequestRequestTypeadminRemove,
		"adminupdate":  AccessPackageAssignmentRequestRequestTypeadminUpdate,
		"notspecified": AccessPackageAssignmentRequestRequestTypenotSpecified,
		"onbehalfadd":  AccessPackageAssignmentRequestRequestTypeonBehalfAdd,
		"systemadd":    AccessPackageAssignmentRequestRequestTypesystemAdd,
		"systemremove": AccessPackageAssignmentRequestRequestTypesystemRemove,
		"systemupdate": AccessPackageAssignmentRequestRequestTypesystemUpdate,
		"useradd":      AccessPackageAssignmentRequestRequestTypeuserAdd,
		"userremove":   AccessPackageAssignmentRequestRequestTypeuserRemove,
		"userupdate":   AccessPackageAssignmentRequestRequestTypeuserUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestRequestType(input)
	return &out, nil
}
