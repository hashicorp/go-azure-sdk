package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      AndroidPkcsCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService           AndroidPkcsCertificateProfileSubjectAlternativeNameType = "domainNameService"
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress                AndroidPkcsCertificateProfileSubjectAlternativeNameType = "emailAddress"
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_None                        AndroidPkcsCertificateProfileSubjectAlternativeNameType = "none"
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier AndroidPkcsCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	AndroidPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName           AndroidPkcsCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForAndroidPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_None),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(AndroidPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *AndroidPkcsCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidPkcsCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidPkcsCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           AndroidPkcsCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                AndroidPkcsCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        AndroidPkcsCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": AndroidPkcsCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           AndroidPkcsCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
