package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileSubjectAlternativeNameType string

const (
	MacOSScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      MacOSScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	MacOSScepCertificateProfileSubjectAlternativeNameTypedomainNameService           MacOSScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	MacOSScepCertificateProfileSubjectAlternativeNameTypeemailAddress                MacOSScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	MacOSScepCertificateProfileSubjectAlternativeNameTypenone                        MacOSScepCertificateProfileSubjectAlternativeNameType = "None"
	MacOSScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier MacOSScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	MacOSScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           MacOSScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForMacOSScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypenone),
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(MacOSScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseMacOSScepCertificateProfileSubjectAlternativeNameType(input string) (*MacOSScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           MacOSScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                MacOSScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        MacOSScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": MacOSScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           MacOSScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
