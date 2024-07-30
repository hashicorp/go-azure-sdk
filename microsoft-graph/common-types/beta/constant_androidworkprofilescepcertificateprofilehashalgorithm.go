package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileHashAlgorithm string

const (
	AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha1 AndroidWorkProfileScepCertificateProfileHashAlgorithm = "sha1"
	AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha2 AndroidWorkProfileScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha1),
		string(AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileHashAlgorithm(input string) (*AndroidWorkProfileScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileHashAlgorithm{
		"sha1": AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": AndroidWorkProfileScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
