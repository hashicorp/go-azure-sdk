package gallerysharingupdate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingProfileGroupTypes string

const (
	SharingProfileGroupTypesAADTenants    SharingProfileGroupTypes = "AADTenants"
	SharingProfileGroupTypesSubscriptions SharingProfileGroupTypes = "Subscriptions"
)

func PossibleValuesForSharingProfileGroupTypes() []string {
	return []string{
		string(SharingProfileGroupTypesAADTenants),
		string(SharingProfileGroupTypesSubscriptions),
	}
}

type SharingUpdateOperationTypes string

const (
	SharingUpdateOperationTypesAdd    SharingUpdateOperationTypes = "Add"
	SharingUpdateOperationTypesRemove SharingUpdateOperationTypes = "Remove"
	SharingUpdateOperationTypesReset  SharingUpdateOperationTypes = "Reset"
)

func PossibleValuesForSharingUpdateOperationTypes() []string {
	return []string{
		string(SharingUpdateOperationTypesAdd),
		string(SharingUpdateOperationTypesRemove),
		string(SharingUpdateOperationTypesReset),
	}
}
