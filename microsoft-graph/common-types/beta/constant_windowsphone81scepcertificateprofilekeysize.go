package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileKeySize string

const (
	WindowsPhone81SCEPCertificateProfileKeySize_Size1024 WindowsPhone81SCEPCertificateProfileKeySize = "size1024"
	WindowsPhone81SCEPCertificateProfileKeySize_Size2048 WindowsPhone81SCEPCertificateProfileKeySize = "size2048"
	WindowsPhone81SCEPCertificateProfileKeySize_Size4096 WindowsPhone81SCEPCertificateProfileKeySize = "size4096"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileKeySize() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileKeySize_Size1024),
		string(WindowsPhone81SCEPCertificateProfileKeySize_Size2048),
		string(WindowsPhone81SCEPCertificateProfileKeySize_Size4096),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileKeySize(input string) (*WindowsPhone81SCEPCertificateProfileKeySize, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileKeySize{
		"size1024": WindowsPhone81SCEPCertificateProfileKeySize_Size1024,
		"size2048": WindowsPhone81SCEPCertificateProfileKeySize_Size2048,
		"size4096": WindowsPhone81SCEPCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileKeySize(input)
	return &out, nil
}
