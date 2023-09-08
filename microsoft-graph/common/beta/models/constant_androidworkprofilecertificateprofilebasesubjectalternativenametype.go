package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypenone                        AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "None"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
