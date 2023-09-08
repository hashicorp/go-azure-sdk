package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseSubjectAlternativeNameType string

const (
	Windows10CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      Windows10CertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	Windows10CertificateProfileBaseSubjectAlternativeNameTypedomainNameService           Windows10CertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	Windows10CertificateProfileBaseSubjectAlternativeNameTypeemailAddress                Windows10CertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	Windows10CertificateProfileBaseSubjectAlternativeNameTypenone                        Windows10CertificateProfileBaseSubjectAlternativeNameType = "None"
	Windows10CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier Windows10CertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	Windows10CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           Windows10CertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindows10CertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypenone),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindows10CertificateProfileBaseSubjectAlternativeNameType(input string) (*Windows10CertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]Windows10CertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      Windows10CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           Windows10CertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                Windows10CertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        Windows10CertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": Windows10CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           Windows10CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
