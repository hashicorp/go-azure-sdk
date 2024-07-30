package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType string

const (
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None                        AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "none"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input string) (*AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
