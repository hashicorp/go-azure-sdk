package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileKeySize string

const (
	AospDeviceOwnerScepCertificateProfileKeySize_Size1024 AospDeviceOwnerScepCertificateProfileKeySize = "size1024"
	AospDeviceOwnerScepCertificateProfileKeySize_Size2048 AospDeviceOwnerScepCertificateProfileKeySize = "size2048"
	AospDeviceOwnerScepCertificateProfileKeySize_Size4096 AospDeviceOwnerScepCertificateProfileKeySize = "size4096"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileKeySize() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileKeySize_Size1024),
		string(AospDeviceOwnerScepCertificateProfileKeySize_Size2048),
		string(AospDeviceOwnerScepCertificateProfileKeySize_Size4096),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileKeySize(input string) (*AospDeviceOwnerScepCertificateProfileKeySize, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileKeySize{
		"size1024": AospDeviceOwnerScepCertificateProfileKeySize_Size1024,
		"size2048": AospDeviceOwnerScepCertificateProfileKeySize_Size2048,
		"size4096": AospDeviceOwnerScepCertificateProfileKeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileKeySize(input)
	return &out, nil
}
