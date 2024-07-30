package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileHashAlgorithm string

const (
	AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1 AospDeviceOwnerScepCertificateProfileHashAlgorithm = "sha1"
	AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2 AospDeviceOwnerScepCertificateProfileHashAlgorithm = "sha2"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1),
		string(AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileHashAlgorithm(input string) (*AospDeviceOwnerScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileHashAlgorithm{
		"sha1": AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha1,
		"sha2": AospDeviceOwnerScepCertificateProfileHashAlgorithm_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
