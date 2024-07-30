package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileHashAlgorithm string

const (
	MacOSScepCertificateProfileHashAlgorithm_Sha1 MacOSScepCertificateProfileHashAlgorithm = "sha1"
	MacOSScepCertificateProfileHashAlgorithm_Sha2 MacOSScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForMacOSScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(MacOSScepCertificateProfileHashAlgorithm_Sha1),
		string(MacOSScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *MacOSScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileHashAlgorithm(input string) (*MacOSScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]MacOSScepCertificateProfileHashAlgorithm{
		"sha1": MacOSScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": MacOSScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
