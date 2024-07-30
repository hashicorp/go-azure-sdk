package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      AndroidCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           AndroidCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	AndroidCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                AndroidCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	AndroidCertificateProfileBaseSubjectAlternativeNameType_None                        AndroidCertificateProfileBaseSubjectAlternativeNameType = "none"
	AndroidCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier AndroidCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           AndroidCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_None),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
