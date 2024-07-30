package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat string

const (
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_CustomAzureADAttribute      ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "customAzureADAttribute"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_DomainNameService           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "domainNameService"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_EmailAddress                ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "emailAddress"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_None                        ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "none"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UniversalResourceIdentifier ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "universalResourceIdentifier"
	ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UserPrincipalName           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat = "userPrincipalName"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_CustomAzureADAttribute),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_DomainNameService),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_EmailAddress),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_None),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UniversalResourceIdentifier),
		string(ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UserPrincipalName),
	}
}

func (s *ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat(input string) (*ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat{
		"customazureadattribute":      ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_CustomAzureADAttribute,
		"domainnameservice":           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_DomainNameService,
		"emailaddress":                ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_EmailAddress,
		"none":                        ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_None,
		"universalresourceidentifier": ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UniversalResourceIdentifier,
		"userprincipalname":           ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateSubjectAlternativeNameFormat(input)
	return &out, nil
}
