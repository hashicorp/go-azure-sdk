package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileKeySize string

const (
	IosScepCertificateProfileKeySize_Size1024 IosScepCertificateProfileKeySize = "size1024"
	IosScepCertificateProfileKeySize_Size2048 IosScepCertificateProfileKeySize = "size2048"
	IosScepCertificateProfileKeySize_Size4096 IosScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForIosScepCertificateProfileKeySize() []string {
	return []string{
		string(IosScepCertificateProfileKeySize_Size1024),
		string(IosScepCertificateProfileKeySize_Size2048),
		string(IosScepCertificateProfileKeySize_Size4096),
	}
}

func (s *IosScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileKeySize(input string) (*IosScepCertificateProfileKeySize, error) {
	vals := map[string]IosScepCertificateProfileKeySize{
		"size1024": IosScepCertificateProfileKeySize_Size1024,
		"size2048": IosScepCertificateProfileKeySize_Size2048,
		"size4096": IosScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileKeySize(input)
	return &out, nil
}
