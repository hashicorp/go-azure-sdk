package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypenone                        AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "None"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidForWorkScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
