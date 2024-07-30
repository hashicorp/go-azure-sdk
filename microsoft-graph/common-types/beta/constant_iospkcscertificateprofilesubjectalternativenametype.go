package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileSubjectAlternativeNameType string

const (
	IosPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      IosPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	IosPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           IosPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	IosPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                IosPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	IosPkcsCertificateProfileSubjectAlternativeNameType_None                        IosPkcsCertificateProfileSubjectAlternativeNameType = "none"
	IosPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier IosPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	IosPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           IosPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForIosPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(IosPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *IosPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosPkcsCertificateProfileSubjectAlternativeNameType(input string) (*IosPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]IosPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      IosPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           IosPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                IosPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        IosPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": IosPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           IosPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
