package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_None                        AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
