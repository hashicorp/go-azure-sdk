package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XCustomSubjectAlternativeNameSanType string

const (
	Windows10XCustomSubjectAlternativeNameSanTypecustomAzureADAttribute      Windows10XCustomSubjectAlternativeNameSanType = "CustomAzureADAttribute"
	Windows10XCustomSubjectAlternativeNameSanTypedomainNameService           Windows10XCustomSubjectAlternativeNameSanType = "DomainNameService"
	Windows10XCustomSubjectAlternativeNameSanTypeemailAddress                Windows10XCustomSubjectAlternativeNameSanType = "EmailAddress"
	Windows10XCustomSubjectAlternativeNameSanTypenone                        Windows10XCustomSubjectAlternativeNameSanType = "None"
	Windows10XCustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier Windows10XCustomSubjectAlternativeNameSanType = "UniversalResourceIdentifier"
	Windows10XCustomSubjectAlternativeNameSanTypeuserPrincipalName           Windows10XCustomSubjectAlternativeNameSanType = "UserPrincipalName"
)

func PossibleValuesForWindows10XCustomSubjectAlternativeNameSanType() []string {
	return []string{
		string(Windows10XCustomSubjectAlternativeNameSanTypecustomAzureADAttribute),
		string(Windows10XCustomSubjectAlternativeNameSanTypedomainNameService),
		string(Windows10XCustomSubjectAlternativeNameSanTypeemailAddress),
		string(Windows10XCustomSubjectAlternativeNameSanTypenone),
		string(Windows10XCustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier),
		string(Windows10XCustomSubjectAlternativeNameSanTypeuserPrincipalName),
	}
}

func parseWindows10XCustomSubjectAlternativeNameSanType(input string) (*Windows10XCustomSubjectAlternativeNameSanType, error) {
	vals := map[string]Windows10XCustomSubjectAlternativeNameSanType{
		"customazureadattribute":      Windows10XCustomSubjectAlternativeNameSanTypecustomAzureADAttribute,
		"domainnameservice":           Windows10XCustomSubjectAlternativeNameSanTypedomainNameService,
		"emailaddress":                Windows10XCustomSubjectAlternativeNameSanTypeemailAddress,
		"none":                        Windows10XCustomSubjectAlternativeNameSanTypenone,
		"universalresourceidentifier": Windows10XCustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier,
		"userprincipalname":           Windows10XCustomSubjectAlternativeNameSanTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XCustomSubjectAlternativeNameSanType(input)
	return &out, nil
}
