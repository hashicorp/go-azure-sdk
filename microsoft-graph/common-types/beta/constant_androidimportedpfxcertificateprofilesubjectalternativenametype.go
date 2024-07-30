package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_None                        AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
