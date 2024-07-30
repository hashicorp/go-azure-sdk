package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_None                        AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "none"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
