package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_None                        WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
