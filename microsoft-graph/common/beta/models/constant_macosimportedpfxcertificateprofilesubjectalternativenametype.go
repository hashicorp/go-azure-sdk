package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseMacOSImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*MacOSImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           MacOSImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
