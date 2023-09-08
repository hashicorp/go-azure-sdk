package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypenone                        AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "None"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
