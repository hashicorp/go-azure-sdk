package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
