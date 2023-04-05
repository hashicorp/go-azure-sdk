package updatesummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthState string

const (
	HealthStateError      HealthState = "Error"
	HealthStateFailure    HealthState = "Failure"
	HealthStateInProgress HealthState = "InProgress"
	HealthStateSuccess    HealthState = "Success"
	HealthStateUnknown    HealthState = "Unknown"
	HealthStateWarning    HealthState = "Warning"
)

func PossibleValuesForHealthState() []string {
	return []string{
		string(HealthStateError),
		string(HealthStateFailure),
		string(HealthStateInProgress),
		string(HealthStateSuccess),
		string(HealthStateUnknown),
		string(HealthStateWarning),
	}
}

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

type Severity string

const (
	SeverityCritical      Severity = "Critical"
	SeverityHidden        Severity = "Hidden"
	SeverityInformational Severity = "Informational"
	SeverityWarning       Severity = "Warning"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityCritical),
		string(SeverityHidden),
		string(SeverityInformational),
		string(SeverityWarning),
	}
}

type Status string

const (
	StatusFailed     Status = "Failed"
	StatusInProgress Status = "InProgress"
	StatusSucceeded  Status = "Succeeded"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusFailed),
		string(StatusInProgress),
		string(StatusSucceeded),
	}
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
