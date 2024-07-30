package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_None                        AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "none"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_None),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
