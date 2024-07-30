package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileKeyUsage string

const (
	IosScepCertificateProfileKeyUsage_DigitalSignature IosScepCertificateProfileKeyUsage = "digitalSignature"
	IosScepCertificateProfileKeyUsage_KeyEncipherment  IosScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForIosScepCertificateProfileKeyUsage() []string {
	return []string{
		string(IosScepCertificateProfileKeyUsage_DigitalSignature),
		string(IosScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *IosScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileKeyUsage(input string) (*IosScepCertificateProfileKeyUsage, error) {
	vals := map[string]IosScepCertificateProfileKeyUsage{
		"digitalsignature": IosScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  IosScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileKeyUsage(input)
	return &out, nil
}
