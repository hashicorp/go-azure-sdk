package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone                        AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "None"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
