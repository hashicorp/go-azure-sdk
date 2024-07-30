package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileHashAlgorithm string

const (
	AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1 AndroidDeviceOwnerScepCertificateProfileHashAlgorithm = "sha1"
	AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2 AndroidDeviceOwnerScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1),
		string(AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileHashAlgorithm(input string) (*AndroidDeviceOwnerScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileHashAlgorithm{
		"sha1": AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": AndroidDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
