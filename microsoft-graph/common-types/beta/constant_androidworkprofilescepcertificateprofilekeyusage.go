package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileKeyUsage string

const (
	AndroidWorkProfileScepCertificateProfileKeyUsage_DigitalSignature AndroidWorkProfileScepCertificateProfileKeyUsage = "digitalSignature"
	AndroidWorkProfileScepCertificateProfileKeyUsage_KeyEncipherment  AndroidWorkProfileScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileKeyUsage_DigitalSignature),
		string(AndroidWorkProfileScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileKeyUsage(input string) (*AndroidWorkProfileScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidWorkProfileScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  AndroidWorkProfileScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileKeyUsage(input)
	return &out, nil
}
