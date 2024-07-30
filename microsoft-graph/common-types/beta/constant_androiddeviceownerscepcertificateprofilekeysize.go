package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileKeySize string

const (
	AndroidDeviceOwnerScepCertificateProfileKeySize_Size1024 AndroidDeviceOwnerScepCertificateProfileKeySize = "size1024"
	AndroidDeviceOwnerScepCertificateProfileKeySize_Size2048 AndroidDeviceOwnerScepCertificateProfileKeySize = "size2048"
	AndroidDeviceOwnerScepCertificateProfileKeySize_Size4096 AndroidDeviceOwnerScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileKeySize_Size1024),
		string(AndroidDeviceOwnerScepCertificateProfileKeySize_Size2048),
		string(AndroidDeviceOwnerScepCertificateProfileKeySize_Size4096),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileKeySize(input string) (*AndroidDeviceOwnerScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileKeySize{
		"size1024": AndroidDeviceOwnerScepCertificateProfileKeySize_Size1024,
		"size2048": AndroidDeviceOwnerScepCertificateProfileKeySize_Size2048,
		"size4096": AndroidDeviceOwnerScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileKeySize(input)
	return &out, nil
}
