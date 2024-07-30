package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileHashAlgorithm string

const (
	Windows81SCEPCertificateProfileHashAlgorithm_Sha1 Windows81SCEPCertificateProfileHashAlgorithm = "sha1"
	Windows81SCEPCertificateProfileHashAlgorithm_Sha2 Windows81SCEPCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForWindows81SCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(Windows81SCEPCertificateProfileHashAlgorithm_Sha1),
		string(Windows81SCEPCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *Windows81SCEPCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileHashAlgorithm(input string) (*Windows81SCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]Windows81SCEPCertificateProfileHashAlgorithm{
		"sha1": Windows81SCEPCertificateProfileHashAlgorithm_Sha1,
		"sha2": Windows81SCEPCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
