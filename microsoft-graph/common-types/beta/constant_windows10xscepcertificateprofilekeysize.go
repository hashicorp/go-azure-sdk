package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileKeySize string

const (
	Windows10XSCEPCertificateProfileKeySize_Size1024 Windows10XSCEPCertificateProfileKeySize = "size1024"
	Windows10XSCEPCertificateProfileKeySize_Size2048 Windows10XSCEPCertificateProfileKeySize = "size2048"
	Windows10XSCEPCertificateProfileKeySize_Size4096 Windows10XSCEPCertificateProfileKeySize = "size4096"
)

func PossibleValuesForWindows10XSCEPCertificateProfileKeySize() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileKeySize_Size1024),
		string(Windows10XSCEPCertificateProfileKeySize_Size2048),
		string(Windows10XSCEPCertificateProfileKeySize_Size4096),
	}
}

func (s *Windows10XSCEPCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileKeySize(input string) (*Windows10XSCEPCertificateProfileKeySize, error) {
	vals := map[string]Windows10XSCEPCertificateProfileKeySize{
		"size1024": Windows10XSCEPCertificateProfileKeySize_Size1024,
		"size2048": Windows10XSCEPCertificateProfileKeySize_Size2048,
		"size4096": Windows10XSCEPCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileKeySize(input)
	return &out, nil
}
