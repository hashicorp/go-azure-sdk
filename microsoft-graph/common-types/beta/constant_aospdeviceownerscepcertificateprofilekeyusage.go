package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileKeyUsage string

const (
	AospDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature AospDeviceOwnerScepCertificateProfileKeyUsage = "digitalSignature"
	AospDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment  AospDeviceOwnerScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature),
		string(AospDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileKeyUsage(input string) (*AospDeviceOwnerScepCertificateProfileKeyUsage, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileKeyUsage{
		"digitalsignature": AospDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  AospDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileKeyUsage(input)
	return &out, nil
}
