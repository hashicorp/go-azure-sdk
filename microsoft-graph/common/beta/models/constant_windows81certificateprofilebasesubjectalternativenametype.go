package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseSubjectAlternativeNameType string

const (
	Windows81CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      Windows81CertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	Windows81CertificateProfileBaseSubjectAlternativeNameTypedomainNameService           Windows81CertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	Windows81CertificateProfileBaseSubjectAlternativeNameTypeemailAddress                Windows81CertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	Windows81CertificateProfileBaseSubjectAlternativeNameTypenone                        Windows81CertificateProfileBaseSubjectAlternativeNameType = "None"
	Windows81CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier Windows81CertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	Windows81CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           Windows81CertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindows81CertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypenone),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindows81CertificateProfileBaseSubjectAlternativeNameType(input string) (*Windows81CertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]Windows81CertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      Windows81CertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           Windows81CertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                Windows81CertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        Windows81CertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": Windows81CertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           Windows81CertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
