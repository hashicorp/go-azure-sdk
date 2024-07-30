package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileKeyUsage string

const (
	WindowsPhone81SCEPCertificateProfileKeyUsage_DigitalSignature WindowsPhone81SCEPCertificateProfileKeyUsage = "digitalSignature"
	WindowsPhone81SCEPCertificateProfileKeyUsage_KeyEncipherment  WindowsPhone81SCEPCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileKeyUsage() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileKeyUsage_DigitalSignature),
		string(WindowsPhone81SCEPCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileKeyUsage(input string) (*WindowsPhone81SCEPCertificateProfileKeyUsage, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileKeyUsage{
		"digitalsignature": WindowsPhone81SCEPCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  WindowsPhone81SCEPCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileKeyUsage(input)
	return &out, nil
}
