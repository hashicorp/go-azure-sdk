package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileKeyUsage string

const (
	AndroidDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature AndroidDeviceOwnerScepCertificateProfileKeyUsage = "digitalSignature"
	AndroidDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment  AndroidDeviceOwnerScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature),
		string(AndroidDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileKeyUsage(input string) (*AndroidDeviceOwnerScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidDeviceOwnerScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  AndroidDeviceOwnerScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileKeyUsage(input)
	return &out, nil
}
