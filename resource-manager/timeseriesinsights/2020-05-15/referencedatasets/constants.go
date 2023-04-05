package referencedatasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataStringComparisonBehavior string

const (
	DataStringComparisonBehaviorOrdinal           DataStringComparisonBehavior = "Ordinal"
	DataStringComparisonBehaviorOrdinalIgnoreCase DataStringComparisonBehavior = "OrdinalIgnoreCase"
)

func PossibleValuesForDataStringComparisonBehavior() []string {
	return []string{
		string(DataStringComparisonBehaviorOrdinal),
		string(DataStringComparisonBehaviorOrdinalIgnoreCase),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type ReferenceDataKeyPropertyType string

const (
	ReferenceDataKeyPropertyTypeBool     ReferenceDataKeyPropertyType = "Bool"
	ReferenceDataKeyPropertyTypeDateTime ReferenceDataKeyPropertyType = "DateTime"
	ReferenceDataKeyPropertyTypeDouble   ReferenceDataKeyPropertyType = "Double"
	ReferenceDataKeyPropertyTypeString   ReferenceDataKeyPropertyType = "String"
)

func PossibleValuesForReferenceDataKeyPropertyType() []string {
	return []string{
		string(ReferenceDataKeyPropertyTypeBool),
		string(ReferenceDataKeyPropertyTypeDateTime),
		string(ReferenceDataKeyPropertyTypeDouble),
		string(ReferenceDataKeyPropertyTypeString),
	}
}
