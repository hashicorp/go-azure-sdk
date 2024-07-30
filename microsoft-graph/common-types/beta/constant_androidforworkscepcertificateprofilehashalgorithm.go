package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileHashAlgorithm string

const (
	AndroidForWorkScepCertificateProfileHashAlgorithm_Sha1 AndroidForWorkScepCertificateProfileHashAlgorithm = "sha1"
	AndroidForWorkScepCertificateProfileHashAlgorithm_Sha2 AndroidForWorkScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileHashAlgorithm_Sha1),
		string(AndroidForWorkScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *AndroidForWorkScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileHashAlgorithm(input string) (*AndroidForWorkScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileHashAlgorithm{
		"sha1": AndroidForWorkScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": AndroidForWorkScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
