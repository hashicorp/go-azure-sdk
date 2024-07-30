package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileKeyUsage string

const (
	MacOSScepCertificateProfileKeyUsage_DigitalSignature MacOSScepCertificateProfileKeyUsage = "digitalSignature"
	MacOSScepCertificateProfileKeyUsage_KeyEncipherment  MacOSScepCertificateProfileKeyUsage = "keyEncipherment"
)

func PossibleValuesForMacOSScepCertificateProfileKeyUsage() []string {
	return []string{
		string(MacOSScepCertificateProfileKeyUsage_DigitalSignature),
		string(MacOSScepCertificateProfileKeyUsage_KeyEncipherment),
	}
}

func (s *MacOSScepCertificateProfileKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileKeyUsage(input string) (*MacOSScepCertificateProfileKeyUsage, error) {
	vals := map[string]MacOSScepCertificateProfileKeyUsage{
		"digitalsignature": MacOSScepCertificateProfileKeyUsage_DigitalSignature,
		"keyencipherment":  MacOSScepCertificateProfileKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileKeyUsage(input)
	return &out, nil
}
