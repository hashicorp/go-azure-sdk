package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequestRequestType string

const (
	AccessPackageResourceRequestRequestTypeadminAdd     AccessPackageResourceRequestRequestType = "AdminAdd"
	AccessPackageResourceRequestRequestTypeadminRemove  AccessPackageResourceRequestRequestType = "AdminRemove"
	AccessPackageResourceRequestRequestTypeadminUpdate  AccessPackageResourceRequestRequestType = "AdminUpdate"
	AccessPackageResourceRequestRequestTypenotSpecified AccessPackageResourceRequestRequestType = "NotSpecified"
	AccessPackageResourceRequestRequestTypeonBehalfAdd  AccessPackageResourceRequestRequestType = "OnBehalfAdd"
	AccessPackageResourceRequestRequestTypesystemAdd    AccessPackageResourceRequestRequestType = "SystemAdd"
	AccessPackageResourceRequestRequestTypesystemRemove AccessPackageResourceRequestRequestType = "SystemRemove"
	AccessPackageResourceRequestRequestTypesystemUpdate AccessPackageResourceRequestRequestType = "SystemUpdate"
	AccessPackageResourceRequestRequestTypeuserAdd      AccessPackageResourceRequestRequestType = "UserAdd"
	AccessPackageResourceRequestRequestTypeuserRemove   AccessPackageResourceRequestRequestType = "UserRemove"
	AccessPackageResourceRequestRequestTypeuserUpdate   AccessPackageResourceRequestRequestType = "UserUpdate"
)

func PossibleValuesForAccessPackageResourceRequestRequestType() []string {
	return []string{
		string(AccessPackageResourceRequestRequestTypeadminAdd),
		string(AccessPackageResourceRequestRequestTypeadminRemove),
		string(AccessPackageResourceRequestRequestTypeadminUpdate),
		string(AccessPackageResourceRequestRequestTypenotSpecified),
		string(AccessPackageResourceRequestRequestTypeonBehalfAdd),
		string(AccessPackageResourceRequestRequestTypesystemAdd),
		string(AccessPackageResourceRequestRequestTypesystemRemove),
		string(AccessPackageResourceRequestRequestTypesystemUpdate),
		string(AccessPackageResourceRequestRequestTypeuserAdd),
		string(AccessPackageResourceRequestRequestTypeuserRemove),
		string(AccessPackageResourceRequestRequestTypeuserUpdate),
	}
}

func parseAccessPackageResourceRequestRequestType(input string) (*AccessPackageResourceRequestRequestType, error) {
	vals := map[string]AccessPackageResourceRequestRequestType{
		"adminadd":     AccessPackageResourceRequestRequestTypeadminAdd,
		"adminremove":  AccessPackageResourceRequestRequestTypeadminRemove,
		"adminupdate":  AccessPackageResourceRequestRequestTypeadminUpdate,
		"notspecified": AccessPackageResourceRequestRequestTypenotSpecified,
		"onbehalfadd":  AccessPackageResourceRequestRequestTypeonBehalfAdd,
		"systemadd":    AccessPackageResourceRequestRequestTypesystemAdd,
		"systemremove": AccessPackageResourceRequestRequestTypesystemRemove,
		"systemupdate": AccessPackageResourceRequestRequestTypesystemUpdate,
		"useradd":      AccessPackageResourceRequestRequestTypeuserAdd,
		"userremove":   AccessPackageResourceRequestRequestTypeuserRemove,
		"userupdate":   AccessPackageResourceRequestRequestTypeuserUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageResourceRequestRequestType(input)
	return &out, nil
}
