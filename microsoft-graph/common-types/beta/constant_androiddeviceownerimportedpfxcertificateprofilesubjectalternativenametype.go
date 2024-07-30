package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_None                        AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
