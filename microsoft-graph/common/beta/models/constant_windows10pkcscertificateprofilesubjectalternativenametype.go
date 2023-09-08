package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileSubjectAlternativeNameType string

const (
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      Windows10PkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           Windows10PkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                Windows10PkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypenone                        Windows10PkcsCertificateProfileSubjectAlternativeNameType = "None"
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier Windows10PkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           Windows10PkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindows10PkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindows10PkcsCertificateProfileSubjectAlternativeNameType(input string) (*Windows10PkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]Windows10PkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      Windows10PkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           Windows10PkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                Windows10PkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        Windows10PkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           Windows10PkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
