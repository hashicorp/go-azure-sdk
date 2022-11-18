package updates

import "strings"

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

func parseAvailabilityType(input string) (*AvailabilityType, error) {
	vals := map[string]AvailabilityType{
		"local":  AvailabilityTypeLocal,
		"notify": AvailabilityTypeNotify,
		"online": AvailabilityTypeOnline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AvailabilityType(input)
	return &out, nil
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

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"downloadfailed":     StateDownloadFailed,
		"downloading":        StateDownloading,
		"hasprerequisite":    StateHasPrerequisite,
		"healthcheckfailed":  StateHealthCheckFailed,
		"healthchecking":     StateHealthChecking,
		"installationfailed": StateInstallationFailed,
		"installed":          StateInstalled,
		"installing":         StateInstalling,
		"invalid":            StateInvalid,
		"notapplicablebecauseanotherupdateisinprogress": StateNotApplicableBecauseAnotherUpdateIsInProgress,
		"obsolete":          StateObsolete,
		"preparationfailed": StatePreparationFailed,
		"preparing":         StatePreparing,
		"ready":             StateReady,
		"readytoinstall":    StateReadyToInstall,
		"recalled":          StateRecalled,
		"scanfailed":        StateScanFailed,
		"scaninprogress":    StateScanInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}
