package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationSigningCertificateType string

const (
	IosEasEmailProfileConfigurationSigningCertificateType_Certificate       IosEasEmailProfileConfigurationSigningCertificateType = "certificate"
	IosEasEmailProfileConfigurationSigningCertificateType_DerivedCredential IosEasEmailProfileConfigurationSigningCertificateType = "derivedCredential"
	IosEasEmailProfileConfigurationSigningCertificateType_None              IosEasEmailProfileConfigurationSigningCertificateType = "none"
)

func PossibleValuesForIosEasEmailProfileConfigurationSigningCertificateType() []string {
	return []string{
		string(IosEasEmailProfileConfigurationSigningCertificateType_Certificate),
		string(IosEasEmailProfileConfigurationSigningCertificateType_DerivedCredential),
		string(IosEasEmailProfileConfigurationSigningCertificateType_None),
	}
}

func (s *IosEasEmailProfileConfigurationSigningCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationSigningCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationSigningCertificateType(input string) (*IosEasEmailProfileConfigurationSigningCertificateType, error) {
	vals := map[string]IosEasEmailProfileConfigurationSigningCertificateType{
		"certificate":       IosEasEmailProfileConfigurationSigningCertificateType_Certificate,
		"derivedcredential": IosEasEmailProfileConfigurationSigningCertificateType_DerivedCredential,
		"none":              IosEasEmailProfileConfigurationSigningCertificateType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationSigningCertificateType(input)
	return &out, nil
}
