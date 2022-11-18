package updatesummaries

import "strings"

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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type UpdateSummariesPropertiesState string

const (
	UpdateSummariesPropertiesStateAppliedSuccessfully   UpdateSummariesPropertiesState = "AppliedSuccessfully"
	UpdateSummariesPropertiesStateNeedsAttention        UpdateSummariesPropertiesState = "NeedsAttention"
	UpdateSummariesPropertiesStatePreparationFailed     UpdateSummariesPropertiesState = "PreparationFailed"
	UpdateSummariesPropertiesStatePreparationInProgress UpdateSummariesPropertiesState = "PreparationInProgress"
	UpdateSummariesPropertiesStateUnknown               UpdateSummariesPropertiesState = "Unknown"
	UpdateSummariesPropertiesStateUpdateAvailable       UpdateSummariesPropertiesState = "UpdateAvailable"
	UpdateSummariesPropertiesStateUpdateFailed          UpdateSummariesPropertiesState = "UpdateFailed"
	UpdateSummariesPropertiesStateUpdateInProgress      UpdateSummariesPropertiesState = "UpdateInProgress"
)

func PossibleValuesForUpdateSummariesPropertiesState() []string {
	return []string{
		string(UpdateSummariesPropertiesStateAppliedSuccessfully),
		string(UpdateSummariesPropertiesStateNeedsAttention),
		string(UpdateSummariesPropertiesStatePreparationFailed),
		string(UpdateSummariesPropertiesStatePreparationInProgress),
		string(UpdateSummariesPropertiesStateUnknown),
		string(UpdateSummariesPropertiesStateUpdateAvailable),
		string(UpdateSummariesPropertiesStateUpdateFailed),
		string(UpdateSummariesPropertiesStateUpdateInProgress),
	}
}

func parseUpdateSummariesPropertiesState(input string) (*UpdateSummariesPropertiesState, error) {
	vals := map[string]UpdateSummariesPropertiesState{
		"appliedsuccessfully":   UpdateSummariesPropertiesStateAppliedSuccessfully,
		"needsattention":        UpdateSummariesPropertiesStateNeedsAttention,
		"preparationfailed":     UpdateSummariesPropertiesStatePreparationFailed,
		"preparationinprogress": UpdateSummariesPropertiesStatePreparationInProgress,
		"unknown":               UpdateSummariesPropertiesStateUnknown,
		"updateavailable":       UpdateSummariesPropertiesStateUpdateAvailable,
		"updatefailed":          UpdateSummariesPropertiesStateUpdateFailed,
		"updateinprogress":      UpdateSummariesPropertiesStateUpdateInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateSummariesPropertiesState(input)
	return &out, nil
}
