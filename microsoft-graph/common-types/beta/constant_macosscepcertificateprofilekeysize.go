package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileKeySize string

const (
	MacOSScepCertificateProfileKeySize_Size1024 MacOSScepCertificateProfileKeySize = "size1024"
	MacOSScepCertificateProfileKeySize_Size2048 MacOSScepCertificateProfileKeySize = "size2048"
	MacOSScepCertificateProfileKeySize_Size4096 MacOSScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForMacOSScepCertificateProfileKeySize() []string {
	return []string{
		string(MacOSScepCertificateProfileKeySize_Size1024),
		string(MacOSScepCertificateProfileKeySize_Size2048),
		string(MacOSScepCertificateProfileKeySize_Size4096),
	}
}

func (s *MacOSScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileKeySize(input string) (*MacOSScepCertificateProfileKeySize, error) {
	vals := map[string]MacOSScepCertificateProfileKeySize{
		"size1024": MacOSScepCertificateProfileKeySize_Size1024,
		"size2048": MacOSScepCertificateProfileKeySize_Size2048,
		"size4096": MacOSScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileKeySize(input)
	return &out, nil
}
