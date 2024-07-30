package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_None                        AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "none"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_None),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
