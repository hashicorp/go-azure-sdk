package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindows10ImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
