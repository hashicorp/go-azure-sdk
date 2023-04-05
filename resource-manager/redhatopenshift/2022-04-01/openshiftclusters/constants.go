package openshiftclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionAtHost string

const (
	EncryptionAtHostDisabled EncryptionAtHost = "Disabled"
	EncryptionAtHostEnabled  EncryptionAtHost = "Enabled"
)

func PossibleValuesForEncryptionAtHost() []string {
	return []string{
		string(EncryptionAtHostDisabled),
		string(EncryptionAtHostEnabled),
	}
}

type FipsValidatedModules string

const (
	FipsValidatedModulesDisabled FipsValidatedModules = "Disabled"
	FipsValidatedModulesEnabled  FipsValidatedModules = "Enabled"
)

func PossibleValuesForFipsValidatedModules() []string {
	return []string{
		string(FipsValidatedModulesDisabled),
		string(FipsValidatedModulesEnabled),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAdminUpdating ProvisioningState = "AdminUpdating"
	ProvisioningStateCreating      ProvisioningState = "Creating"
	ProvisioningStateDeleting      ProvisioningState = "Deleting"
	ProvisioningStateFailed        ProvisioningState = "Failed"
	ProvisioningStateSucceeded     ProvisioningState = "Succeeded"
	ProvisioningStateUpdating      ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAdminUpdating),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type Visibility string

const (
	VisibilityPrivate Visibility = "Private"
	VisibilityPublic  Visibility = "Public"
)

func PossibleValuesForVisibility() []string {
	return []string{
		string(VisibilityPrivate),
		string(VisibilityPublic),
	}
}
