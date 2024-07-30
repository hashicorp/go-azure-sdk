package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEncryptionCertificateType string

const (
	IosEasEmailProfileConfigurationEncryptionCertificateType_Certificate       IosEasEmailProfileConfigurationEncryptionCertificateType = "certificate"
	IosEasEmailProfileConfigurationEncryptionCertificateType_DerivedCredential IosEasEmailProfileConfigurationEncryptionCertificateType = "derivedCredential"
	IosEasEmailProfileConfigurationEncryptionCertificateType_None              IosEasEmailProfileConfigurationEncryptionCertificateType = "none"
)

func PossibleValuesForIosEasEmailProfileConfigurationEncryptionCertificateType() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEncryptionCertificateType_Certificate),
		string(IosEasEmailProfileConfigurationEncryptionCertificateType_DerivedCredential),
		string(IosEasEmailProfileConfigurationEncryptionCertificateType_None),
	}
}

func (s *IosEasEmailProfileConfigurationEncryptionCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationEncryptionCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationEncryptionCertificateType(input string) (*IosEasEmailProfileConfigurationEncryptionCertificateType, error) {
	vals := map[string]IosEasEmailProfileConfigurationEncryptionCertificateType{
		"certificate":       IosEasEmailProfileConfigurationEncryptionCertificateType_Certificate,
		"derivedcredential": IosEasEmailProfileConfigurationEncryptionCertificateType_DerivedCredential,
		"none":              IosEasEmailProfileConfigurationEncryptionCertificateType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEncryptionCertificateType(input)
	return &out, nil
}
