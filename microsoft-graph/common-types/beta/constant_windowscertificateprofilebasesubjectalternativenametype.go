package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseSubjectAlternativeNameType string

const (
	WindowsCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      WindowsCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	WindowsCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           WindowsCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	WindowsCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                WindowsCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	WindowsCertificateProfileBaseSubjectAlternativeNameType_None                        WindowsCertificateProfileBaseSubjectAlternativeNameType = "none"
	WindowsCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier WindowsCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	WindowsCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           WindowsCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindowsCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_None),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *WindowsCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsCertificateProfileBaseSubjectAlternativeNameType(input string) (*WindowsCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]WindowsCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      WindowsCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           WindowsCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                WindowsCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        WindowsCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": WindowsCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           WindowsCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
