package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None                        AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "none"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
