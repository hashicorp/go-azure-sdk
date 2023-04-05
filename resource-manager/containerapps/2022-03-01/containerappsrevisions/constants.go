package containerappsrevisions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevisionHealthState string

const (
	RevisionHealthStateHealthy   RevisionHealthState = "Healthy"
	RevisionHealthStateNone      RevisionHealthState = "None"
	RevisionHealthStateUnhealthy RevisionHealthState = "Unhealthy"
)

func PossibleValuesForRevisionHealthState() []string {
	return []string{
		string(RevisionHealthStateHealthy),
		string(RevisionHealthStateNone),
		string(RevisionHealthStateUnhealthy),
	}
}

type RevisionProvisioningState string

const (
	RevisionProvisioningStateDeprovisioned  RevisionProvisioningState = "Deprovisioned"
	RevisionProvisioningStateDeprovisioning RevisionProvisioningState = "Deprovisioning"
	RevisionProvisioningStateFailed         RevisionProvisioningState = "Failed"
	RevisionProvisioningStateProvisioned    RevisionProvisioningState = "Provisioned"
	RevisionProvisioningStateProvisioning   RevisionProvisioningState = "Provisioning"
)

func PossibleValuesForRevisionProvisioningState() []string {
	return []string{
		string(RevisionProvisioningStateDeprovisioned),
		string(RevisionProvisioningStateDeprovisioning),
		string(RevisionProvisioningStateFailed),
		string(RevisionProvisioningStateProvisioned),
		string(RevisionProvisioningStateProvisioning),
	}
}

type Scheme string

const (
	SchemeHTTP  Scheme = "HTTP"
	SchemeHTTPS Scheme = "HTTPS"
)

func PossibleValuesForScheme() []string {
	return []string{
		string(SchemeHTTP),
		string(SchemeHTTPS),
	}
}

type StorageType string

const (
	StorageTypeAzureFile StorageType = "AzureFile"
	StorageTypeEmptyDir  StorageType = "EmptyDir"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypeAzureFile),
		string(StorageTypeEmptyDir),
	}
}

type Type string

const (
	TypeLiveness  Type = "Liveness"
	TypeReadiness Type = "Readiness"
	TypeStartup   Type = "Startup"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLiveness),
		string(TypeReadiness),
		string(TypeStartup),
	}
}
