package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileHashAlgorithm string

const (
	WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha1 WindowsPhone81SCEPCertificateProfileHashAlgorithm = "sha1"
	WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha2 WindowsPhone81SCEPCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha1),
		string(WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileHashAlgorithm(input string) (*WindowsPhone81SCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileHashAlgorithm{
		"sha1": WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha1,
		"sha2": WindowsPhone81SCEPCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
