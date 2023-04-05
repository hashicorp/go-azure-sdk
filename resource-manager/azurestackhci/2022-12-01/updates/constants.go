package updates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityType string

const (
	AvailabilityTypeLocal  AvailabilityType = "Local"
	AvailabilityTypeNotify AvailabilityType = "Notify"
	AvailabilityTypeOnline AvailabilityType = "Online"
)

func PossibleValuesForAvailabilityType() []string {
	return []string{
		string(AvailabilityTypeLocal),
		string(AvailabilityTypeNotify),
		string(AvailabilityTypeOnline),
	}
}

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

type RebootRequirement string

const (
	RebootRequirementFalse   RebootRequirement = "False"
	RebootRequirementTrue    RebootRequirement = "True"
	RebootRequirementUnknown RebootRequirement = "Unknown"
)

func PossibleValuesForRebootRequirement() []string {
	return []string{
		string(RebootRequirementFalse),
		string(RebootRequirementTrue),
		string(RebootRequirementUnknown),
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

type State string

const (
	StateDownloadFailed                                State = "DownloadFailed"
	StateDownloading                                   State = "Downloading"
	StateHasPrerequisite                               State = "HasPrerequisite"
	StateHealthCheckFailed                             State = "HealthCheckFailed"
	StateHealthChecking                                State = "HealthChecking"
	StateInstallationFailed                            State = "InstallationFailed"
	StateInstalled                                     State = "Installed"
	StateInstalling                                    State = "Installing"
	StateInvalid                                       State = "Invalid"
	StateNotApplicableBecauseAnotherUpdateIsInProgress State = "NotApplicableBecauseAnotherUpdateIsInProgress"
	StateObsolete                                      State = "Obsolete"
	StatePreparationFailed                             State = "PreparationFailed"
	StatePreparing                                     State = "Preparing"
	StateReady                                         State = "Ready"
	StateReadyToInstall                                State = "ReadyToInstall"
	StateRecalled                                      State = "Recalled"
	StateScanFailed                                    State = "ScanFailed"
	StateScanInProgress                                State = "ScanInProgress"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateDownloadFailed),
		string(StateDownloading),
		string(StateHasPrerequisite),
		string(StateHealthCheckFailed),
		string(StateHealthChecking),
		string(StateInstallationFailed),
		string(StateInstalled),
		string(StateInstalling),
		string(StateInvalid),
		string(StateNotApplicableBecauseAnotherUpdateIsInProgress),
		string(StateObsolete),
		string(StatePreparationFailed),
		string(StatePreparing),
		string(StateReady),
		string(StateReadyToInstall),
		string(StateRecalled),
		string(StateScanFailed),
		string(StateScanInProgress),
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
