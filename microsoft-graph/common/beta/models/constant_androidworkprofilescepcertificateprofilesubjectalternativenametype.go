package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypenone                        AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "None"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
