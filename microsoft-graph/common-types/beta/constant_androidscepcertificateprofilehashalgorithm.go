package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileHashAlgorithm string

const (
	AndroidScepCertificateProfileHashAlgorithm_Sha1 AndroidScepCertificateProfileHashAlgorithm = "sha1"
	AndroidScepCertificateProfileHashAlgorithm_Sha2 AndroidScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForAndroidScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidScepCertificateProfileHashAlgorithm_Sha1),
		string(AndroidScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *AndroidScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileHashAlgorithm(input string) (*AndroidScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidScepCertificateProfileHashAlgorithm{
		"sha1": AndroidScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": AndroidScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
