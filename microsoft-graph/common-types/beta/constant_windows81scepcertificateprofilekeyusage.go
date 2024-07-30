package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileKeyUsage string

const (
	Windows81SCEPCertificateProfileKeyUsage_DigitalSignature Windows81SCEPCertificateProfileKeyUsage = "digitalSignature"
	Windows81SCEPCertificateProfileKeyUsage_KeyEncipherment  Windows81SCEPCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForWindows81SCEPCertificateProfileKeyUsage() []string {
	return []string{
		string(Windows81SCEPCertificateProfileKeyUsage_DigitalSignature),
		string(Windows81SCEPCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *Windows81SCEPCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileKeyUsage(input string) (*Windows81SCEPCertificateProfileKeyUsage, error) {
	vals := map[string]Windows81SCEPCertificateProfileKeyUsage{
		"digitalsignature": Windows81SCEPCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  Windows81SCEPCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileKeyUsage(input)
	return &out, nil
}
