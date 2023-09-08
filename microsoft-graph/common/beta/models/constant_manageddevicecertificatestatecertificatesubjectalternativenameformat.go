package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat string

const (
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatcustomAzureADAttribute      ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "CustomAzureADAttribute"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatdomainNameService           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "DomainNameService"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatemailAddress                ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "EmailAddress"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatnone                        ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "None"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuniversalResourceIdentifier ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "UniversalResourceIdentifier"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuserPrincipalName           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "UserPrincipalName"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatcustomAzureADAttribute),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatdomainNameService),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatemailAddress),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatnone),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuniversalResourceIdentifier),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuserPrincipalName),
	}
}

func parseManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat(input string) (*ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat{
		"customazureadattribute":      ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatcustomAzureADAttribute,
		"domainnameservice":           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatdomainNameService,
		"emailaddress":                ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatemailAddress,
		"none":                        ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatnone,
		"universalresourceidentifier": ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuniversalResourceIdentifier,
		"userprincipalname":           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormatuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat(input)
	return &out, nil
}
