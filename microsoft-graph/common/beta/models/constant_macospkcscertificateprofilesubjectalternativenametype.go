package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileSubjectAlternativeNameType string

const (
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      MacOSPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           MacOSPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                MacOSPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypenone                        MacOSPkcsCertificateProfileSubjectAlternativeNameType = "None"
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier MacOSPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           MacOSPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForMacOSPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseMacOSPkcsCertificateProfileSubjectAlternativeNameType(input string) (*MacOSPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           MacOSPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                MacOSPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        MacOSPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           MacOSPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
