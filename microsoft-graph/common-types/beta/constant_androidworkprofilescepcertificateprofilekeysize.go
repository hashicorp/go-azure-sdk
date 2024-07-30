package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileKeySize string

const (
	AndroidWorkProfileScepCertificateProfileKeySize_Size1024 AndroidWorkProfileScepCertificateProfileKeySize = "size1024"
	AndroidWorkProfileScepCertificateProfileKeySize_Size2048 AndroidWorkProfileScepCertificateProfileKeySize = "size2048"
	AndroidWorkProfileScepCertificateProfileKeySize_Size4096 AndroidWorkProfileScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileKeySize_Size1024),
		string(AndroidWorkProfileScepCertificateProfileKeySize_Size2048),
		string(AndroidWorkProfileScepCertificateProfileKeySize_Size4096),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileKeySize(input string) (*AndroidWorkProfileScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileKeySize{
		"size1024": AndroidWorkProfileScepCertificateProfileKeySize_Size1024,
		"size2048": AndroidWorkProfileScepCertificateProfileKeySize_Size2048,
		"size4096": AndroidWorkProfileScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileKeySize(input)
	return &out, nil
}
