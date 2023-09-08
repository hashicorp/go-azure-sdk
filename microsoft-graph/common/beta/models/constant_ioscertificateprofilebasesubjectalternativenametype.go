package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseSubjectAlternativeNameType string

const (
	IosCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      IosCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	IosCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           IosCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	IosCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                IosCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	IosCertificateProfileBaseSubjectAlternativeNameTypenone                        IosCertificateProfileBaseSubjectAlternativeNameType = "None"
	IosCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier IosCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	IosCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           IosCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForIosCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(IosCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(IosCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(IosCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(IosCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(IosCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(IosCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseIosCertificateProfileBaseSubjectAlternativeNameType(input string) (*IosCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]IosCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      IosCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           IosCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                IosCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        IosCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": IosCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           IosCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
