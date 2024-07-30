package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None                        AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "none"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
