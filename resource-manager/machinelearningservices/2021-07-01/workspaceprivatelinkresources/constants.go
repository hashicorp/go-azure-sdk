package workspaceprivatelinkresources

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
)

func PossibleValuesForResourceIdentityType() []string {
	return []string{
		string(ResourceIdentityTypeNone),
		string(ResourceIdentityTypeSystemAssigned),
		string(ResourceIdentityTypeSystemAssignedUserAssigned),
		string(ResourceIdentityTypeUserAssigned),
	}
}

func parseResourceIdentityType(input string) (*ResourceIdentityType, error) {
	vals := map[string]ResourceIdentityType{
		"none":                        ResourceIdentityTypeNone,
		"systemassigned":              ResourceIdentityTypeSystemAssigned,
		"systemassigned,userassigned": ResourceIdentityTypeSystemAssignedUserAssigned,
		"userassigned":                ResourceIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceIdentityType(input)
	return &out, nil
}
