package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileKeySize string

const (
	AndroidForWorkScepCertificateProfileKeySize_Size1024 AndroidForWorkScepCertificateProfileKeySize = "size1024"
	AndroidForWorkScepCertificateProfileKeySize_Size2048 AndroidForWorkScepCertificateProfileKeySize = "size2048"
	AndroidForWorkScepCertificateProfileKeySize_Size4096 AndroidForWorkScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileKeySize_Size1024),
		string(AndroidForWorkScepCertificateProfileKeySize_Size2048),
		string(AndroidForWorkScepCertificateProfileKeySize_Size4096),
	}
}

func (s *AndroidForWorkScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileKeySize(input string) (*AndroidForWorkScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileKeySize{
		"size1024": AndroidForWorkScepCertificateProfileKeySize_Size1024,
		"size2048": AndroidForWorkScepCertificateProfileKeySize_Size2048,
		"size4096": AndroidForWorkScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileKeySize(input)
	return &out, nil
}
