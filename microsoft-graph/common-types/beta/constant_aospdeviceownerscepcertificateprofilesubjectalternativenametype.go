package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType string

const (
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress                AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None                        AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "none"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input string) (*AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
