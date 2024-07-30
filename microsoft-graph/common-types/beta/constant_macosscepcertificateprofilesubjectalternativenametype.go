package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileSubjectAlternativeNameType string

const (
	MacOSScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      MacOSScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	MacOSScepCertificateProfileSubjectAlternativeNameType_DomainNameService           MacOSScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	MacOSScepCertificateProfileSubjectAlternativeNameType_EmailAddress                MacOSScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	MacOSScepCertificateProfileSubjectAlternativeNameType_None                        MacOSScepCertificateProfileSubjectAlternativeNameType = "none"
	MacOSScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier MacOSScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	MacOSScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           MacOSScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForMacOSScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_None),
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(MacOSScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *MacOSScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileSubjectAlternativeNameType(input string) (*MacOSScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           MacOSScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                MacOSScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        MacOSScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": MacOSScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           MacOSScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
