package servicerunners

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                       ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned             ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	ManagedIdentityTypeUserAssigned               ManagedIdentityType = "UserAssigned"
)

func PossibleValuesForManagedIdentityType() []string {
	return []string{
		string(ManagedIdentityTypeNone),
		string(ManagedIdentityTypeSystemAssigned),
		string(ManagedIdentityTypeSystemAssignedUserAssigned),
		string(ManagedIdentityTypeUserAssigned),
	}
}
