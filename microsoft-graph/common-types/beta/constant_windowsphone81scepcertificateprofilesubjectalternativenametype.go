package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType string

const (
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "domainNameService"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress                WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "emailAddress"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_None                        WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "none"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_None),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType(input string) (*WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
