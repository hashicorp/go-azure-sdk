package afdcustomdomains

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AfdCertificateType string

const (
	AfdCertificateTypeAzureFirstPartyManagedCertificate AfdCertificateType = "AzureFirstPartyManagedCertificate"
	AfdCertificateTypeCustomerCertificate               AfdCertificateType = "CustomerCertificate"
	AfdCertificateTypeManagedCertificate                AfdCertificateType = "ManagedCertificate"
)

func PossibleValuesForAfdCertificateType() []string {
	return []string{
		string(AfdCertificateTypeAzureFirstPartyManagedCertificate),
		string(AfdCertificateTypeCustomerCertificate),
		string(AfdCertificateTypeManagedCertificate),
	}
}

func (s *AfdCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAfdCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAfdCertificateType(input string) (*AfdCertificateType, error) {
	vals := map[string]AfdCertificateType{
		"azurefirstpartymanagedcertificate": AfdCertificateTypeAzureFirstPartyManagedCertificate,
		"customercertificate":               AfdCertificateTypeCustomerCertificate,
		"managedcertificate":                AfdCertificateTypeManagedCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AfdCertificateType(input)
	return &out, nil
}

type AfdMinimumTlsVersion string

const (
	AfdMinimumTlsVersionTLSOneTwo  AfdMinimumTlsVersion = "TLS12"
	AfdMinimumTlsVersionTLSOneZero AfdMinimumTlsVersion = "TLS10"
)

func PossibleValuesForAfdMinimumTlsVersion() []string {
	return []string{
		string(AfdMinimumTlsVersionTLSOneTwo),
		string(AfdMinimumTlsVersionTLSOneZero),
	}
}

func (s *AfdMinimumTlsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAfdMinimumTlsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAfdMinimumTlsVersion(input string) (*AfdMinimumTlsVersion, error) {
	vals := map[string]AfdMinimumTlsVersion{
		"tls12": AfdMinimumTlsVersionTLSOneTwo,
		"tls10": AfdMinimumTlsVersionTLSOneZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AfdMinimumTlsVersion(input)
	return &out, nil
}

type AfdProvisioningState string

const (
	AfdProvisioningStateCreating  AfdProvisioningState = "Creating"
	AfdProvisioningStateDeleting  AfdProvisioningState = "Deleting"
	AfdProvisioningStateFailed    AfdProvisioningState = "Failed"
	AfdProvisioningStateSucceeded AfdProvisioningState = "Succeeded"
	AfdProvisioningStateUpdating  AfdProvisioningState = "Updating"
)

func PossibleValuesForAfdProvisioningState() []string {
	return []string{
		string(AfdProvisioningStateCreating),
		string(AfdProvisioningStateDeleting),
		string(AfdProvisioningStateFailed),
		string(AfdProvisioningStateSucceeded),
		string(AfdProvisioningStateUpdating),
	}
}

func (s *AfdProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAfdProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAfdProvisioningState(input string) (*AfdProvisioningState, error) {
	vals := map[string]AfdProvisioningState{
		"creating":  AfdProvisioningStateCreating,
		"deleting":  AfdProvisioningStateDeleting,
		"failed":    AfdProvisioningStateFailed,
		"succeeded": AfdProvisioningStateSucceeded,
		"updating":  AfdProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AfdProvisioningState(input)
	return &out, nil
}

type DeploymentStatus string

const (
	DeploymentStatusFailed     DeploymentStatus = "Failed"
	DeploymentStatusInProgress DeploymentStatus = "InProgress"
	DeploymentStatusNotStarted DeploymentStatus = "NotStarted"
	DeploymentStatusSucceeded  DeploymentStatus = "Succeeded"
)

func PossibleValuesForDeploymentStatus() []string {
	return []string{
		string(DeploymentStatusFailed),
		string(DeploymentStatusInProgress),
		string(DeploymentStatusNotStarted),
		string(DeploymentStatusSucceeded),
	}
}

func (s *DeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentStatus(input string) (*DeploymentStatus, error) {
	vals := map[string]DeploymentStatus{
		"failed":     DeploymentStatusFailed,
		"inprogress": DeploymentStatusInProgress,
		"notstarted": DeploymentStatusNotStarted,
		"succeeded":  DeploymentStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentStatus(input)
	return &out, nil
}

type DomainValidationState string

const (
	DomainValidationStateApproved                  DomainValidationState = "Approved"
	DomainValidationStateInternalError             DomainValidationState = "InternalError"
	DomainValidationStatePending                   DomainValidationState = "Pending"
	DomainValidationStatePendingRevalidation       DomainValidationState = "PendingRevalidation"
	DomainValidationStateRefreshingValidationToken DomainValidationState = "RefreshingValidationToken"
	DomainValidationStateRejected                  DomainValidationState = "Rejected"
	DomainValidationStateSubmitting                DomainValidationState = "Submitting"
	DomainValidationStateTimedOut                  DomainValidationState = "TimedOut"
	DomainValidationStateUnknown                   DomainValidationState = "Unknown"
)

func PossibleValuesForDomainValidationState() []string {
	return []string{
		string(DomainValidationStateApproved),
		string(DomainValidationStateInternalError),
		string(DomainValidationStatePending),
		string(DomainValidationStatePendingRevalidation),
		string(DomainValidationStateRefreshingValidationToken),
		string(DomainValidationStateRejected),
		string(DomainValidationStateSubmitting),
		string(DomainValidationStateTimedOut),
		string(DomainValidationStateUnknown),
	}
}

func (s *DomainValidationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDomainValidationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDomainValidationState(input string) (*DomainValidationState, error) {
	vals := map[string]DomainValidationState{
		"approved":                  DomainValidationStateApproved,
		"internalerror":             DomainValidationStateInternalError,
		"pending":                   DomainValidationStatePending,
		"pendingrevalidation":       DomainValidationStatePendingRevalidation,
		"refreshingvalidationtoken": DomainValidationStateRefreshingValidationToken,
		"rejected":                  DomainValidationStateRejected,
		"submitting":                DomainValidationStateSubmitting,
		"timedout":                  DomainValidationStateTimedOut,
		"unknown":                   DomainValidationStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DomainValidationState(input)
	return &out, nil
}
