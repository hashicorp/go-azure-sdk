package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SymantecCodeSigningCertificateStatus string

const (
	SymantecCodeSigningCertificateStatus_NotProvisioned SymantecCodeSigningCertificateStatus = "notProvisioned"
	SymantecCodeSigningCertificateStatus_Provisioned    SymantecCodeSigningCertificateStatus = "provisioned"
)

func PossibleValuesForSymantecCodeSigningCertificateStatus() []string {
	return []string{
		string(SymantecCodeSigningCertificateStatus_NotProvisioned),
		string(SymantecCodeSigningCertificateStatus_Provisioned),
	}
}

func (s *SymantecCodeSigningCertificateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSymantecCodeSigningCertificateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSymantecCodeSigningCertificateStatus(input string) (*SymantecCodeSigningCertificateStatus, error) {
	vals := map[string]SymantecCodeSigningCertificateStatus{
		"notprovisioned": SymantecCodeSigningCertificateStatus_NotProvisioned,
		"provisioned":    SymantecCodeSigningCertificateStatus_Provisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SymantecCodeSigningCertificateStatus(input)
	return &out, nil
}
