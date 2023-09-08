package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypenone                        AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "None"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
