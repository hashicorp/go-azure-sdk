package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSubjectAlternativeNameSanType string

const (
	CustomSubjectAlternativeNameSanTypecustomAzureADAttribute      CustomSubjectAlternativeNameSanType = "CustomAzureADAttribute"
	CustomSubjectAlternativeNameSanTypedomainNameService           CustomSubjectAlternativeNameSanType = "DomainNameService"
	CustomSubjectAlternativeNameSanTypeemailAddress                CustomSubjectAlternativeNameSanType = "EmailAddress"
	CustomSubjectAlternativeNameSanTypenone                        CustomSubjectAlternativeNameSanType = "None"
	CustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier CustomSubjectAlternativeNameSanType = "UniversalResourceIdentifier"
	CustomSubjectAlternativeNameSanTypeuserPrincipalName           CustomSubjectAlternativeNameSanType = "UserPrincipalName"
)

func PossibleValuesForCustomSubjectAlternativeNameSanType() []string {
	return []string{
		string(CustomSubjectAlternativeNameSanTypecustomAzureADAttribute),
		string(CustomSubjectAlternativeNameSanTypedomainNameService),
		string(CustomSubjectAlternativeNameSanTypeemailAddress),
		string(CustomSubjectAlternativeNameSanTypenone),
		string(CustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier),
		string(CustomSubjectAlternativeNameSanTypeuserPrincipalName),
	}
}

func parseCustomSubjectAlternativeNameSanType(input string) (*CustomSubjectAlternativeNameSanType, error) {
	vals := map[string]CustomSubjectAlternativeNameSanType{
		"customazureadattribute":      CustomSubjectAlternativeNameSanTypecustomAzureADAttribute,
		"domainnameservice":           CustomSubjectAlternativeNameSanTypedomainNameService,
		"emailaddress":                CustomSubjectAlternativeNameSanTypeemailAddress,
		"none":                        CustomSubjectAlternativeNameSanTypenone,
		"universalresourceidentifier": CustomSubjectAlternativeNameSanTypeuniversalResourceIdentifier,
		"userprincipalname":           CustomSubjectAlternativeNameSanTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomSubjectAlternativeNameSanType(input)
	return &out, nil
}
