package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileSubjectAlternativeNameType string

const (
	IosScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      IosScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	IosScepCertificateProfileSubjectAlternativeNameTypedomainNameService           IosScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	IosScepCertificateProfileSubjectAlternativeNameTypeemailAddress                IosScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	IosScepCertificateProfileSubjectAlternativeNameTypenone                        IosScepCertificateProfileSubjectAlternativeNameType = "None"
	IosScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier IosScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	IosScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           IosScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForIosScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(IosScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(IosScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(IosScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(IosScepCertificateProfileSubjectAlternativeNameTypenone),
		string(IosScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(IosScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseIosScepCertificateProfileSubjectAlternativeNameType(input string) (*IosScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]IosScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      IosScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           IosScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                IosScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        IosScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": IosScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           IosScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
