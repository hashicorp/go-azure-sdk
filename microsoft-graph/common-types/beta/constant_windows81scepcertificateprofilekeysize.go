package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileKeySize string

const (
	Windows81SCEPCertificateProfileKeySize_Size1024 Windows81SCEPCertificateProfileKeySize = "size1024"
	Windows81SCEPCertificateProfileKeySize_Size2048 Windows81SCEPCertificateProfileKeySize = "size2048"
	Windows81SCEPCertificateProfileKeySize_Size4096 Windows81SCEPCertificateProfileKeySize = "size4096"
)

func PossibleValuesForWindows81SCEPCertificateProfileKeySize() []string {
	return []string{
		string(Windows81SCEPCertificateProfileKeySize_Size1024),
		string(Windows81SCEPCertificateProfileKeySize_Size2048),
		string(Windows81SCEPCertificateProfileKeySize_Size4096),
	}
}

func (s *Windows81SCEPCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileKeySize(input string) (*Windows81SCEPCertificateProfileKeySize, error) {
	vals := map[string]Windows81SCEPCertificateProfileKeySize{
		"size1024": Windows81SCEPCertificateProfileKeySize_Size1024,
		"size2048": Windows81SCEPCertificateProfileKeySize_Size2048,
		"size4096": Windows81SCEPCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileKeySize(input)
	return &out, nil
}
