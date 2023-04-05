package packetcorecontrolplanes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeAAD      AuthenticationType = "AAD"
	AuthenticationTypePassword AuthenticationType = "Password"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeAAD),
		string(AuthenticationTypePassword),
	}
}

type BillingSku string

const (
	BillingSkuGFive    BillingSku = "G5"
	BillingSkuGOne     BillingSku = "G1"
	BillingSkuGOneZero BillingSku = "G10"
	BillingSkuGTwo     BillingSku = "G2"
	BillingSkuGZero    BillingSku = "G0"
)

func PossibleValuesForBillingSku() []string {
	return []string{
		string(BillingSkuGFive),
		string(BillingSkuGOne),
		string(BillingSkuGOneZero),
		string(BillingSkuGTwo),
		string(BillingSkuGZero),
	}
}

type CertificateProvisioningState string

const (
	CertificateProvisioningStateFailed         CertificateProvisioningState = "Failed"
	CertificateProvisioningStateNotProvisioned CertificateProvisioningState = "NotProvisioned"
	CertificateProvisioningStateProvisioned    CertificateProvisioningState = "Provisioned"
)

func PossibleValuesForCertificateProvisioningState() []string {
	return []string{
		string(CertificateProvisioningStateFailed),
		string(CertificateProvisioningStateNotProvisioned),
		string(CertificateProvisioningStateProvisioned),
	}
}

type CoreNetworkType string

const (
	CoreNetworkTypeEPC    CoreNetworkType = "EPC"
	CoreNetworkTypeFiveGC CoreNetworkType = "5GC"
)

func PossibleValuesForCoreNetworkType() []string {
	return []string{
		string(CoreNetworkTypeEPC),
		string(CoreNetworkTypeFiveGC),
	}
}

type InstallationState string

const (
	InstallationStateFailed       InstallationState = "Failed"
	InstallationStateInstalled    InstallationState = "Installed"
	InstallationStateInstalling   InstallationState = "Installing"
	InstallationStateReinstalling InstallationState = "Reinstalling"
	InstallationStateRollingBack  InstallationState = "RollingBack"
	InstallationStateUninstalled  InstallationState = "Uninstalled"
	InstallationStateUninstalling InstallationState = "Uninstalling"
	InstallationStateUpdating     InstallationState = "Updating"
	InstallationStateUpgrading    InstallationState = "Upgrading"
)

func PossibleValuesForInstallationState() []string {
	return []string{
		string(InstallationStateFailed),
		string(InstallationStateInstalled),
		string(InstallationStateInstalling),
		string(InstallationStateReinstalling),
		string(InstallationStateRollingBack),
		string(InstallationStateUninstalled),
		string(InstallationStateUninstalling),
		string(InstallationStateUpdating),
		string(InstallationStateUpgrading),
	}
}

type PlatformType string

const (
	PlatformTypeAKSNegativeHCI                              PlatformType = "AKS-HCI"
	PlatformTypeThreePNegativeAZURENegativeSTACKNegativeHCI PlatformType = "3P-AZURE-STACK-HCI"
)

func PossibleValuesForPlatformType() []string {
	return []string{
		string(PlatformTypeAKSNegativeHCI),
		string(PlatformTypeThreePNegativeAZURENegativeSTACKNegativeHCI),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
	}
}
