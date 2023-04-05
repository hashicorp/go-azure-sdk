package fluidrelayservers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CmkIdentityType string

const (
	CmkIdentityTypeSystemAssigned CmkIdentityType = "SystemAssigned"
	CmkIdentityTypeUserAssigned   CmkIdentityType = "UserAssigned"
)

func PossibleValuesForCmkIdentityType() []string {
	return []string{
		string(CmkIdentityTypeSystemAssigned),
		string(CmkIdentityTypeUserAssigned),
	}
}

type KeyName string

const (
	KeyNameKeyOne KeyName = "key1"
	KeyNameKeyTwo KeyName = "key2"
)

func PossibleValuesForKeyName() []string {
	return []string{
		string(KeyNameKeyOne),
		string(KeyNameKeyTwo),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

type StorageSKU string

const (
	StorageSKUBasic    StorageSKU = "basic"
	StorageSKUStandard StorageSKU = "standard"
)

func PossibleValuesForStorageSKU() []string {
	return []string{
		string(StorageSKUBasic),
		string(StorageSKUStandard),
	}
}
