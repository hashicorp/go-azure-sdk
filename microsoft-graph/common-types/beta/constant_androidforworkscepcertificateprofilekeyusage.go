package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileKeyUsage string

const (
	AndroidForWorkScepCertificateProfileKeyUsage_DigitalSignature AndroidForWorkScepCertificateProfileKeyUsage = "digitalSignature"
	AndroidForWorkScepCertificateProfileKeyUsage_KeyEncipherment  AndroidForWorkScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileKeyUsage_DigitalSignature),
		string(AndroidForWorkScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *AndroidForWorkScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileKeyUsage(input string) (*AndroidForWorkScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidForWorkScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  AndroidForWorkScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileKeyUsage(input)
	return &out, nil
}
