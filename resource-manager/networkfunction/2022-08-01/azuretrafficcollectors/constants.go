package azuretrafficcollectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

func PossibleValuesForCreatedByType() []string {
	return []string{
		string(CreatedByTypeApplication),
		string(CreatedByTypeKey),
		string(CreatedByTypeManagedIdentity),
		string(CreatedByTypeUser),
	}
}

type DestinationType string

const (
	DestinationTypeAzureMonitor DestinationType = "AzureMonitor"
)

func PossibleValuesForDestinationType() []string {
	return []string{
		string(DestinationTypeAzureMonitor),
	}
}

type EmissionType string

const (
	EmissionTypeIPFIX EmissionType = "IPFIX"
)

func PossibleValuesForEmissionType() []string {
	return []string{
		string(EmissionTypeIPFIX),
	}
}

type IngestionType string

const (
	IngestionTypeIPFIX IngestionType = "IPFIX"
)

func PossibleValuesForIngestionType() []string {
	return []string{
		string(IngestionTypeIPFIX),
	}
}

type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SourceType string

const (
	SourceTypeResource SourceType = "Resource"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeResource),
	}
}
