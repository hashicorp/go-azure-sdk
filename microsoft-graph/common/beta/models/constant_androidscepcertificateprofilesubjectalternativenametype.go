package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidScepCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidScepCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidScepCertificateProfileSubjectAlternativeNameTypenone                        AndroidScepCertificateProfileSubjectAlternativeNameType = "None"
	AndroidScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
