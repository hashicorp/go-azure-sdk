package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None                        AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "none"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
