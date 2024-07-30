package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileSubjectAlternativeNameType string

const (
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      Windows10PkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           Windows10PkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                Windows10PkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_None                        Windows10PkcsCertificateProfileSubjectAlternativeNameType = "none"
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier Windows10PkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	Windows10PkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           Windows10PkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindows10PkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_None),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(Windows10PkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *Windows10PkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PkcsCertificateProfileSubjectAlternativeNameType(input string) (*Windows10PkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]Windows10PkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      Windows10PkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           Windows10PkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                Windows10PkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        Windows10PkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": Windows10PkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           Windows10PkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
