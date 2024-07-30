package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseSubjectAlternativeNameType string

const (
	IosCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      IosCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	IosCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           IosCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	IosCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                IosCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	IosCertificateProfileBaseSubjectAlternativeNameType_None                        IosCertificateProfileBaseSubjectAlternativeNameType = "none"
	IosCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier IosCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	IosCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           IosCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForIosCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(IosCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(IosCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(IosCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(IosCertificateProfileBaseSubjectAlternativeNameType_None),
		string(IosCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(IosCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *IosCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCertificateProfileBaseSubjectAlternativeNameType(input string) (*IosCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]IosCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      IosCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           IosCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                IosCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        IosCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": IosCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           IosCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
