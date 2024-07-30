package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseSubjectAlternativeNameType string

const (
	Windows10CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      Windows10CertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	Windows10CertificateProfileBaseSubjectAlternativeNameType_DomainNameService           Windows10CertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	Windows10CertificateProfileBaseSubjectAlternativeNameType_EmailAddress                Windows10CertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	Windows10CertificateProfileBaseSubjectAlternativeNameType_None                        Windows10CertificateProfileBaseSubjectAlternativeNameType = "none"
	Windows10CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier Windows10CertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	Windows10CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           Windows10CertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindows10CertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_None),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(Windows10CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *Windows10CertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CertificateProfileBaseSubjectAlternativeNameType(input string) (*Windows10CertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]Windows10CertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      Windows10CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           Windows10CertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                Windows10CertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        Windows10CertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": Windows10CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           Windows10CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
