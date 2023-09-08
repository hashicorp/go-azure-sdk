package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
