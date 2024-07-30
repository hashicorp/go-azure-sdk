package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_None                        AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "none"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
