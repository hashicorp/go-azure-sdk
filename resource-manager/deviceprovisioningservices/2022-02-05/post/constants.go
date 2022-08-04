package post

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRightsDescription string

const (
	AccessRightsDescriptionDeviceConnect           AccessRightsDescription = "DeviceConnect"
	AccessRightsDescriptionEnrollmentRead          AccessRightsDescription = "EnrollmentRead"
	AccessRightsDescriptionEnrollmentWrite         AccessRightsDescription = "EnrollmentWrite"
	AccessRightsDescriptionRegistrationStatusRead  AccessRightsDescription = "RegistrationStatusRead"
	AccessRightsDescriptionRegistrationStatusWrite AccessRightsDescription = "RegistrationStatusWrite"
	AccessRightsDescriptionServiceConfig           AccessRightsDescription = "ServiceConfig"
)

func PossibleValuesForAccessRightsDescription() []string {
	return []string{
		string(AccessRightsDescriptionDeviceConnect),
		string(AccessRightsDescriptionEnrollmentRead),
		string(AccessRightsDescriptionEnrollmentWrite),
		string(AccessRightsDescriptionRegistrationStatusRead),
		string(AccessRightsDescriptionRegistrationStatusWrite),
		string(AccessRightsDescriptionServiceConfig),
	}
}

func parseAccessRightsDescription(input string) (*AccessRightsDescription, error) {
	vals := map[string]AccessRightsDescription{
		"deviceconnect":           AccessRightsDescriptionDeviceConnect,
		"enrollmentread":          AccessRightsDescriptionEnrollmentRead,
		"enrollmentwrite":         AccessRightsDescriptionEnrollmentWrite,
		"registrationstatusread":  AccessRightsDescriptionRegistrationStatusRead,
		"registrationstatuswrite": AccessRightsDescriptionRegistrationStatusWrite,
		"serviceconfig":           AccessRightsDescriptionServiceConfig,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessRightsDescription(input)
	return &out, nil
}

type CertificatePurpose string

const (
	CertificatePurposeClientAuthentication CertificatePurpose = "clientAuthentication"
	CertificatePurposeServerAuthentication CertificatePurpose = "serverAuthentication"
)

func PossibleValuesForCertificatePurpose() []string {
	return []string{
		string(CertificatePurposeClientAuthentication),
		string(CertificatePurposeServerAuthentication),
	}
}

func parseCertificatePurpose(input string) (*CertificatePurpose, error) {
	vals := map[string]CertificatePurpose{
		"clientauthentication": CertificatePurposeClientAuthentication,
		"serverauthentication": CertificatePurposeServerAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificatePurpose(input)
	return &out, nil
}

type NameUnavailabilityReason string

const (
	NameUnavailabilityReasonAlreadyExists NameUnavailabilityReason = "AlreadyExists"
	NameUnavailabilityReasonInvalid       NameUnavailabilityReason = "Invalid"
)

func PossibleValuesForNameUnavailabilityReason() []string {
	return []string{
		string(NameUnavailabilityReasonAlreadyExists),
		string(NameUnavailabilityReasonInvalid),
	}
}

func parseNameUnavailabilityReason(input string) (*NameUnavailabilityReason, error) {
	vals := map[string]NameUnavailabilityReason{
		"alreadyexists": NameUnavailabilityReasonAlreadyExists,
		"invalid":       NameUnavailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameUnavailabilityReason(input)
	return &out, nil
}
