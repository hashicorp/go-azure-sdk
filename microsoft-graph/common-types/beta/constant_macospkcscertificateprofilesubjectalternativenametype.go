package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileSubjectAlternativeNameType string

const (
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      MacOSPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           MacOSPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                MacOSPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_None                        MacOSPkcsCertificateProfileSubjectAlternativeNameType = "none"
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier MacOSPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	MacOSPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           MacOSPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForMacOSPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(MacOSPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *MacOSPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPkcsCertificateProfileSubjectAlternativeNameType(input string) (*MacOSPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           MacOSPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                MacOSPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        MacOSPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": MacOSPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           MacOSPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
