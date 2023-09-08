package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseSubjectAlternativeNameType string

const (
	MacOSCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      MacOSCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	MacOSCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           MacOSCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	MacOSCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                MacOSCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	MacOSCertificateProfileBaseSubjectAlternativeNameTypenone                        MacOSCertificateProfileBaseSubjectAlternativeNameType = "None"
	MacOSCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier MacOSCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	MacOSCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           MacOSCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForMacOSCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseMacOSCertificateProfileBaseSubjectAlternativeNameType(input string) (*MacOSCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]MacOSCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      MacOSCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           MacOSCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                MacOSCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        MacOSCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": MacOSCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           MacOSCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
