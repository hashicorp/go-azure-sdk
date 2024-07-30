package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "domainNameService"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress                MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "emailAddress"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_None                        MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "none"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_None),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *MacOSImportedPFXCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSImportedPFXCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*MacOSImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           MacOSImportedPFXCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
