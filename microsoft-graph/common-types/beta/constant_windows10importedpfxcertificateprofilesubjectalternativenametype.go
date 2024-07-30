package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_None                        Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10ImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10ImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
