package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileKeyUsage string

const (
	Windows10XSCEPCertificateProfileKeyUsage_DigitalSignature Windows10XSCEPCertificateProfileKeyUsage = "digitalSignature"
	Windows10XSCEPCertificateProfileKeyUsage_KeyEncipherment  Windows10XSCEPCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForWindows10XSCEPCertificateProfileKeyUsage() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileKeyUsage_DigitalSignature),
		string(Windows10XSCEPCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *Windows10XSCEPCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileKeyUsage(input string) (*Windows10XSCEPCertificateProfileKeyUsage, error) {
	vals := map[string]Windows10XSCEPCertificateProfileKeyUsage{
		"digitalsignature": Windows10XSCEPCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  Windows10XSCEPCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileKeyUsage(input)
	return &out, nil
}
