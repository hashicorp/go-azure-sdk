package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseCodeSigningCertificateStatus string

const (
	EnterpriseCodeSigningCertificateStatus_NotProvisioned EnterpriseCodeSigningCertificateStatus = "notProvisioned"
	EnterpriseCodeSigningCertificateStatus_Provisioned    EnterpriseCodeSigningCertificateStatus = "provisioned"
)

func PossibleValuesForEnterpriseCodeSigningCertificateStatus() []string {
	return []string{
		string(EnterpriseCodeSigningCertificateStatus_NotProvisioned),
		string(EnterpriseCodeSigningCertificateStatus_Provisioned),
	}
}

func (s *EnterpriseCodeSigningCertificateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnterpriseCodeSigningCertificateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnterpriseCodeSigningCertificateStatus(input string) (*EnterpriseCodeSigningCertificateStatus, error) {
	vals := map[string]EnterpriseCodeSigningCertificateStatus{
		"notprovisioned": EnterpriseCodeSigningCertificateStatus_NotProvisioned,
		"provisioned":    EnterpriseCodeSigningCertificateStatus_Provisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnterpriseCodeSigningCertificateStatus(input)
	return &out, nil
}
