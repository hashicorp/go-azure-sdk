package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseSubjectAlternativeNameType string

const (
	Windows81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      Windows81CertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	Windows81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService           Windows81CertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	Windows81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress                Windows81CertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	Windows81CertificateProfileBaseSubjectAlternativeNameType_None                        Windows81CertificateProfileBaseSubjectAlternativeNameType = "none"
	Windows81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier Windows81CertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	Windows81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           Windows81CertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindows81CertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_None),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(Windows81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *Windows81CertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81CertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81CertificateProfileBaseSubjectAlternativeNameType(input string) (*Windows81CertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]Windows81CertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      Windows81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           Windows81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                Windows81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        Windows81CertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": Windows81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           Windows81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
