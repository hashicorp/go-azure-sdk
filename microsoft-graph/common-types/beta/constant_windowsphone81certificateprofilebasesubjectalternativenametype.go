package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType string

const (
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService           WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress                WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_None                        WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "none"
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_None),
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81CertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81CertificateProfileBaseSubjectAlternativeNameType(input string) (*WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
