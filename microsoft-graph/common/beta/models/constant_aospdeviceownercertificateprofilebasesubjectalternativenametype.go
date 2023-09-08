package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType string

const (
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone                        AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "None"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input string) (*AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
