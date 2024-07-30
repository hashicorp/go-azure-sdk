package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None                        AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "none"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
