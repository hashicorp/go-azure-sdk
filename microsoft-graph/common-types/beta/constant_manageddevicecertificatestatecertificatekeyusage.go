package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateKeyUsage string

const (
	ManagedDeviceCertificateStateCertificateKeyUsage_DigitalSignature ManagedDeviceCertificateStateCertificateKeyUsage = "digitalSignature"
	ManagedDeviceCertificateStateCertificateKeyUsage_KeyEncipherment  ManagedDeviceCertificateStateCertificateKeyUsage = "keyEncipherment"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateKeyUsage() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateKeyUsage_DigitalSignature),
		string(ManagedDeviceCertificateStateCertificateKeyUsage_KeyEncipherment),
	}
}

func (s *ManagedDeviceCertificateStateCertificateKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateKeyUsage(input string) (*ManagedDeviceCertificateStateCertificateKeyUsage, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateKeyUsage{
		"digitalsignature": ManagedDeviceCertificateStateCertificateKeyUsage_DigitalSignature,
		"keyencipherment":  ManagedDeviceCertificateStateCertificateKeyUsage_KeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateKeyUsage(input)
	return &out, nil
}
