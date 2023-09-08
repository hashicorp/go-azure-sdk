package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileSubjectAlternativeNameType string

const (
	IosPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      IosPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	IosPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           IosPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	IosPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                IosPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	IosPkcsCertificateProfileSubjectAlternativeNameTypenone                        IosPkcsCertificateProfileSubjectAlternativeNameType = "None"
	IosPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier IosPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	IosPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           IosPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForIosPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(IosPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseIosPkcsCertificateProfileSubjectAlternativeNameType(input string) (*IosPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]IosPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      IosPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           IosPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                IosPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        IosPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": IosPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           IosPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
