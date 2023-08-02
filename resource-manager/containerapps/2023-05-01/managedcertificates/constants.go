package managedcertificates

import "strings"

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

func parseCertificateProvisioningState(input string) (*CertificateProvisioningState, error) {
	vals := map[string]CertificateProvisioningState{
		"canceled":     CertificateProvisioningStateCanceled,
		"deletefailed": CertificateProvisioningStateDeleteFailed,
		"failed":       CertificateProvisioningStateFailed,
		"pending":      CertificateProvisioningStatePending,
		"succeeded":    CertificateProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateProvisioningState(input)
	return &out, nil
}

type ManagedCertificateDomainControlValidation string

const (
	ManagedCertificateDomainControlValidationCNAME ManagedCertificateDomainControlValidation = "CNAME"
	ManagedCertificateDomainControlValidationHTTP  ManagedCertificateDomainControlValidation = "HTTP"
	ManagedCertificateDomainControlValidationTXT   ManagedCertificateDomainControlValidation = "TXT"
)

func PossibleValuesForManagedCertificateDomainControlValidation() []string {
	return []string{
		string(ManagedCertificateDomainControlValidationCNAME),
		string(ManagedCertificateDomainControlValidationHTTP),
		string(ManagedCertificateDomainControlValidationTXT),
	}
}

func parseManagedCertificateDomainControlValidation(input string) (*ManagedCertificateDomainControlValidation, error) {
	vals := map[string]ManagedCertificateDomainControlValidation{
		"cname": ManagedCertificateDomainControlValidationCNAME,
		"http":  ManagedCertificateDomainControlValidationHTTP,
		"txt":   ManagedCertificateDomainControlValidationTXT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedCertificateDomainControlValidation(input)
	return &out, nil
}
