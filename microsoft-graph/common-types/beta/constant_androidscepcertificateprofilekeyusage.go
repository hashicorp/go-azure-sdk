package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileKeyUsage string

const (
	AndroidScepCertificateProfileKeyUsage_DigitalSignature AndroidScepCertificateProfileKeyUsage = "digitalSignature"
	AndroidScepCertificateProfileKeyUsage_KeyEncipherment  AndroidScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForAndroidScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidScepCertificateProfileKeyUsage_DigitalSignature),
		string(AndroidScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *AndroidScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileKeyUsage(input string) (*AndroidScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  AndroidScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileKeyUsage(input)
	return &out, nil
}
