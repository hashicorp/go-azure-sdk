package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XCustomSubjectAlternativeNameSanType string

const (
	Windows10XCustomSubjectAlternativeNameSanType_CustomAzureADAttribute      Windows10XCustomSubjectAlternativeNameSanType = "customAzureADAttribute"
	Windows10XCustomSubjectAlternativeNameSanType_DomainNameService           Windows10XCustomSubjectAlternativeNameSanType = "domainNameService"
	Windows10XCustomSubjectAlternativeNameSanType_EmailAddress                Windows10XCustomSubjectAlternativeNameSanType = "emailAddress"
	Windows10XCustomSubjectAlternativeNameSanType_None                        Windows10XCustomSubjectAlternativeNameSanType = "none"
	Windows10XCustomSubjectAlternativeNameSanType_UniversalResourceIdentifier Windows10XCustomSubjectAlternativeNameSanType = "universalResourceIdentifier"
	Windows10XCustomSubjectAlternativeNameSanType_UserPrincipalName           Windows10XCustomSubjectAlternativeNameSanType = "userPrincipalName"
)

func PossibleValuesForWindows10XCustomSubjectAlternativeNameSanType() []string {
	return []string{
		string(Windows10XCustomSubjectAlternativeNameSanType_CustomAzureADAttribute),
		string(Windows10XCustomSubjectAlternativeNameSanType_DomainNameService),
		string(Windows10XCustomSubjectAlternativeNameSanType_EmailAddress),
		string(Windows10XCustomSubjectAlternativeNameSanType_None),
		string(Windows10XCustomSubjectAlternativeNameSanType_UniversalResourceIdentifier),
		string(Windows10XCustomSubjectAlternativeNameSanType_UserPrincipalName),
	}
}

func (s *Windows10XCustomSubjectAlternativeNameSanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XCustomSubjectAlternativeNameSanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XCustomSubjectAlternativeNameSanType(input string) (*Windows10XCustomSubjectAlternativeNameSanType, error) {
	vals := map[string]Windows10XCustomSubjectAlternativeNameSanType{
		"customazureadattribute":      Windows10XCustomSubjectAlternativeNameSanType_CustomAzureADAttribute,
		"domainnameservice":           Windows10XCustomSubjectAlternativeNameSanType_DomainNameService,
		"emailaddress":                Windows10XCustomSubjectAlternativeNameSanType_EmailAddress,
		"none":                        Windows10XCustomSubjectAlternativeNameSanType_None,
		"universalresourceidentifier": Windows10XCustomSubjectAlternativeNameSanType_UniversalResourceIdentifier,
		"userprincipalname":           Windows10XCustomSubjectAlternativeNameSanType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XCustomSubjectAlternativeNameSanType(input)
	return &out, nil
}
