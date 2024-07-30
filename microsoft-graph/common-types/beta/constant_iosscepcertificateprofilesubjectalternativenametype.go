package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileSubjectAlternativeNameType string

const (
	IosScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute      IosScepCertificateProfileSubjectAlternativeNameType = "customAzureADAttribute"
	IosScepCertificateProfileSubjectAlternativeNameType_DomainNameService           IosScepCertificateProfileSubjectAlternativeNameType = "domainNameService"
	IosScepCertificateProfileSubjectAlternativeNameType_EmailAddress                IosScepCertificateProfileSubjectAlternativeNameType = "emailAddress"
	IosScepCertificateProfileSubjectAlternativeNameType_None                        IosScepCertificateProfileSubjectAlternativeNameType = "none"
	IosScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier IosScepCertificateProfileSubjectAlternativeNameType = "universalResourceIdentifier"
	IosScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName           IosScepCertificateProfileSubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForIosScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(IosScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute),
		string(IosScepCertificateProfileSubjectAlternativeNameType_DomainNameService),
		string(IosScepCertificateProfileSubjectAlternativeNameType_EmailAddress),
		string(IosScepCertificateProfileSubjectAlternativeNameType_None),
		string(IosScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier),
		string(IosScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *IosScepCertificateProfileSubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileSubjectAlternativeNameType(input string) (*IosScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]IosScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      IosScepCertificateProfileSubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           IosScepCertificateProfileSubjectAlternativeNameType_DomainNameService,
		"emailaddress":                IosScepCertificateProfileSubjectAlternativeNameType_EmailAddress,
		"none":                        IosScepCertificateProfileSubjectAlternativeNameType_None,
		"universalresourceidentifier": IosScepCertificateProfileSubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           IosScepCertificateProfileSubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
