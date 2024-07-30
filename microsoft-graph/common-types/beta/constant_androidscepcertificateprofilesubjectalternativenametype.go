package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidScepCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidScepCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidScepCertificateProfileSubjectAlternativeNameType_None                        AndroidScepCertificateProfileSubjectAlternativeNameType = "none"
	AndroidScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
