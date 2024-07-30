package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_None                        AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "none"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidForWorkScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
