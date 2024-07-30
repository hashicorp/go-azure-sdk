package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSubjectAlternativeNameSanType string

const (
	CustomSubjectAlternativeNameSanType_CustomAzureADAttribute      CustomSubjectAlternativeNameSanType = "customAzureADAttribute"
	CustomSubjectAlternativeNameSanType_DomainNameService           CustomSubjectAlternativeNameSanType = "domainNameService"
	CustomSubjectAlternativeNameSanType_EmailAddress                CustomSubjectAlternativeNameSanType = "emailAddress"
	CustomSubjectAlternativeNameSanType_None                        CustomSubjectAlternativeNameSanType = "none"
	CustomSubjectAlternativeNameSanType_UniversalResourceIdentifier CustomSubjectAlternativeNameSanType = "universalResourceIdentifier"
	CustomSubjectAlternativeNameSanType_UserPrincipalName           CustomSubjectAlternativeNameSanType = "userPrincipalName"
)

func PossibleValuesForCustomSubjectAlternativeNameSanType() []string {
	return []string{
		string(CustomSubjectAlternativeNameSanType_CustomAzureADAttribute),
		string(CustomSubjectAlternativeNameSanType_DomainNameService),
		string(CustomSubjectAlternativeNameSanType_EmailAddress),
		string(CustomSubjectAlternativeNameSanType_None),
		string(CustomSubjectAlternativeNameSanType_UniversalResourceIdentifier),
		string(CustomSubjectAlternativeNameSanType_UserPrincipalName),
	}
}

func (s *CustomSubjectAlternativeNameSanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomSubjectAlternativeNameSanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomSubjectAlternativeNameSanType(input string) (*CustomSubjectAlternativeNameSanType, error) {
	vals := map[string]CustomSubjectAlternativeNameSanType{
		"customazureadattribute":      CustomSubjectAlternativeNameSanType_CustomAzureADAttribute,
		"domainnameservice":           CustomSubjectAlternativeNameSanType_DomainNameService,
		"emailaddress":                CustomSubjectAlternativeNameSanType_EmailAddress,
		"none":                        CustomSubjectAlternativeNameSanType_None,
		"universalresourceidentifier": CustomSubjectAlternativeNameSanType_UniversalResourceIdentifier,
		"userprincipalname":           CustomSubjectAlternativeNameSanType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomSubjectAlternativeNameSanType(input)
	return &out, nil
}
