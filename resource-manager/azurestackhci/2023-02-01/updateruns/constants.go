package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
	}
}

type UpdateRunPropertiesState string

const (
	UpdateRunPropertiesStateFailed     UpdateRunPropertiesState = "Failed"
	UpdateRunPropertiesStateInProgress UpdateRunPropertiesState = "InProgress"
	UpdateRunPropertiesStateSucceeded  UpdateRunPropertiesState = "Succeeded"
	UpdateRunPropertiesStateUnknown    UpdateRunPropertiesState = "Unknown"
)

func PossibleValuesForUpdateRunPropertiesState() []string {
	return []string{
		string(UpdateRunPropertiesStateFailed),
		string(UpdateRunPropertiesStateInProgress),
		string(UpdateRunPropertiesStateSucceeded),
		string(UpdateRunPropertiesStateUnknown),
	}
}
