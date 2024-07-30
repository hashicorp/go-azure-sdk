package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_None                        AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "none"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
