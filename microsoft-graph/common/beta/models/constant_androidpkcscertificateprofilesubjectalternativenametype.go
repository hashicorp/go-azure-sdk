package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypenone                        AndroidPkcsCertificateProfileSubjectAlternativeNameType = "None"
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
