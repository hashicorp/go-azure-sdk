package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType string

const (
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypedomainNameService           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeemailAddress                WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypenone                        WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "None"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypenone),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType(input string) (*WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
