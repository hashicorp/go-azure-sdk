package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateProvisioningState string

const (
	CertificateProvisioningStateCanceled     CertificateProvisioningState = "Canceled"
	CertificateProvisioningStateDeleteFailed CertificateProvisioningState = "DeleteFailed"
	CertificateProvisioningStateFailed       CertificateProvisioningState = "Failed"
	CertificateProvisioningStatePending      CertificateProvisioningState = "Pending"
	CertificateProvisioningStateSucceeded    CertificateProvisioningState = "Succeeded"
)

func PossibleValuesForCertificateProvisioningState() []string {
	return []string{
		string(CertificateProvisioningStateCanceled),
		string(CertificateProvisioningStateDeleteFailed),
		string(CertificateProvisioningStateFailed),
		string(CertificateProvisioningStatePending),
		string(CertificateProvisioningStateSucceeded),
	}
}

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

type EnvironmentProvisioningState string

const (
	EnvironmentProvisioningStateCanceled                      EnvironmentProvisioningState = "Canceled"
	EnvironmentProvisioningStateFailed                        EnvironmentProvisioningState = "Failed"
	EnvironmentProvisioningStateInfrastructureSetupComplete   EnvironmentProvisioningState = "InfrastructureSetupComplete"
	EnvironmentProvisioningStateInfrastructureSetupInProgress EnvironmentProvisioningState = "InfrastructureSetupInProgress"
	EnvironmentProvisioningStateInitializationInProgress      EnvironmentProvisioningState = "InitializationInProgress"
	EnvironmentProvisioningStateScheduledForDelete            EnvironmentProvisioningState = "ScheduledForDelete"
	EnvironmentProvisioningStateSucceeded                     EnvironmentProvisioningState = "Succeeded"
	EnvironmentProvisioningStateUpgradeFailed                 EnvironmentProvisioningState = "UpgradeFailed"
	EnvironmentProvisioningStateUpgradeRequested              EnvironmentProvisioningState = "UpgradeRequested"
	EnvironmentProvisioningStateWaiting                       EnvironmentProvisioningState = "Waiting"
)

func PossibleValuesForEnvironmentProvisioningState() []string {
	return []string{
		string(EnvironmentProvisioningStateCanceled),
		string(EnvironmentProvisioningStateFailed),
		string(EnvironmentProvisioningStateInfrastructureSetupComplete),
		string(EnvironmentProvisioningStateInfrastructureSetupInProgress),
		string(EnvironmentProvisioningStateInitializationInProgress),
		string(EnvironmentProvisioningStateScheduledForDelete),
		string(EnvironmentProvisioningStateSucceeded),
		string(EnvironmentProvisioningStateUpgradeFailed),
		string(EnvironmentProvisioningStateUpgradeRequested),
		string(EnvironmentProvisioningStateWaiting),
	}
}
