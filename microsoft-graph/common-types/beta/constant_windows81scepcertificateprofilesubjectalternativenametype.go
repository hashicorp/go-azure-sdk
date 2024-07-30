package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileSubjectAlternativeNameType string

const (
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      Windows81SCEPCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService           Windows81SCEPCertificateProfileSubjectAlternativeNameType = "domainNameService"
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress                Windows81SCEPCertificateProfileSubjectAlternativeNameType = "emailAddress"
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_None                        Windows81SCEPCertificateProfileSubjectAlternativeNameType = "none"
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier Windows81SCEPCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	Windows81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName           Windows81SCEPCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForWindows81SCEPCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_None),
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(Windows81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *Windows81SCEPCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileSubjectAlternativeNameType(input string) (*Windows81SCEPCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]Windows81SCEPCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      Windows81SCEPCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           Windows81SCEPCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                Windows81SCEPCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        Windows81SCEPCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": Windows81SCEPCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           Windows81SCEPCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
