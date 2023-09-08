package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
