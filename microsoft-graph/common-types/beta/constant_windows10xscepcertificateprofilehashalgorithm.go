package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileHashAlgorithm string

const (
	Windows10XSCEPCertificateProfileHashAlgorithm_Sha1 Windows10XSCEPCertificateProfileHashAlgorithm = "sha1"
	Windows10XSCEPCertificateProfileHashAlgorithm_Sha2 Windows10XSCEPCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForWindows10XSCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileHashAlgorithm_Sha1),
		string(Windows10XSCEPCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *Windows10XSCEPCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileHashAlgorithm(input string) (*Windows10XSCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]Windows10XSCEPCertificateProfileHashAlgorithm{
		"sha1": Windows10XSCEPCertificateProfileHashAlgorithm_Sha1,
		"sha2": Windows10XSCEPCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
