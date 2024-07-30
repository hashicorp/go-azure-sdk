package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileKeySize string

const (
	AndroidScepCertificateProfileKeySize_Size1024 AndroidScepCertificateProfileKeySize = "size1024"
	AndroidScepCertificateProfileKeySize_Size2048 AndroidScepCertificateProfileKeySize = "size2048"
	AndroidScepCertificateProfileKeySize_Size4096 AndroidScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForAndroidScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidScepCertificateProfileKeySize_Size1024),
		string(AndroidScepCertificateProfileKeySize_Size2048),
		string(AndroidScepCertificateProfileKeySize_Size4096),
	}
}

func (s *AndroidScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileKeySize(input string) (*AndroidScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidScepCertificateProfileKeySize{
		"size1024": AndroidScepCertificateProfileKeySize_Size1024,
		"size2048": AndroidScepCertificateProfileKeySize_Size2048,
		"size4096": AndroidScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileKeySize(input)
	return &out, nil
}
