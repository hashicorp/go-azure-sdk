package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType string

const (
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress                AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone                        AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "None"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input string) (*AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
