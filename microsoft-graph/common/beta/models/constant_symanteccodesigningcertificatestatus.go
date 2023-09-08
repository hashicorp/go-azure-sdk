package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SymantecCodeSigningCertificateStatus string

const (
	SymantecCodeSigningCertificateStatusnotProvisioned SymantecCodeSigningCertificateStatus = "NotProvisioned"
	SymantecCodeSigningCertificateStatusprovisioned    SymantecCodeSigningCertificateStatus = "Provisioned"
)

func PossibleValuesForSymantecCodeSigningCertificateStatus() []string {
	return []string{
		string(SymantecCodeSigningCertificateStatusnotProvisioned),
		string(SymantecCodeSigningCertificateStatusprovisioned),
	}
}

func parseSymantecCodeSigningCertificateStatus(input string) (*SymantecCodeSigningCertificateStatus, error) {
	vals := map[string]SymantecCodeSigningCertificateStatus{
		"notprovisioned": SymantecCodeSigningCertificateStatusnotProvisioned,
		"provisioned":    SymantecCodeSigningCertificateStatusprovisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SymantecCodeSigningCertificateStatus(input)
	return &out, nil
}
