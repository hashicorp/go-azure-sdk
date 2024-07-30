package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseSubjectAlternativeNameType string

const (
	MacOSCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute      MacOSCertificateProfileBaseSubjectAlternativeNameType = "customAzureADAttribute"
	MacOSCertificateProfileBaseSubjectAlternativeNameType_DomainNameService           MacOSCertificateProfileBaseSubjectAlternativeNameType = "domainNameService"
	MacOSCertificateProfileBaseSubjectAlternativeNameType_EmailAddress                MacOSCertificateProfileBaseSubjectAlternativeNameType = "emailAddress"
	MacOSCertificateProfileBaseSubjectAlternativeNameType_None                        MacOSCertificateProfileBaseSubjectAlternativeNameType = "none"
	MacOSCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier MacOSCertificateProfileBaseSubjectAlternativeNameType = "universalResourceIdentifier"
	MacOSCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName           MacOSCertificateProfileBaseSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForMacOSCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_DomainNameService),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_EmailAddress),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_None),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(MacOSCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *MacOSCertificateProfileBaseSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCertificateProfileBaseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCertificateProfileBaseSubjectAlternativeNameType(input string) (*MacOSCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]MacOSCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      MacOSCertificateProfileBaseSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           MacOSCertificateProfileBaseSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                MacOSCertificateProfileBaseSubjectAlternativeNameType_EmailAddress,
		"none":                        MacOSCertificateProfileBaseSubjectAlternativeNameType_None,
		"universalresourceidentifier": MacOSCertificateProfileBaseSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           MacOSCertificateProfileBaseSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
