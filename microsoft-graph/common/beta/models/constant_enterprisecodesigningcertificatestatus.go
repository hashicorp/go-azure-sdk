package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseCodeSigningCertificateStatus string

const (
	EnterpriseCodeSigningCertificateStatusnotProvisioned EnterpriseCodeSigningCertificateStatus = "NotProvisioned"
	EnterpriseCodeSigningCertificateStatusprovisioned    EnterpriseCodeSigningCertificateStatus = "Provisioned"
)

func PossibleValuesForEnterpriseCodeSigningCertificateStatus() []string {
	return []string{
		string(EnterpriseCodeSigningCertificateStatusnotProvisioned),
		string(EnterpriseCodeSigningCertificateStatusprovisioned),
	}
}

func parseEnterpriseCodeSigningCertificateStatus(input string) (*EnterpriseCodeSigningCertificateStatus, error) {
	vals := map[string]EnterpriseCodeSigningCertificateStatus{
		"notprovisioned": EnterpriseCodeSigningCertificateStatusnotProvisioned,
		"provisioned":    EnterpriseCodeSigningCertificateStatusprovisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnterpriseCodeSigningCertificateStatus(input)
	return &out, nil
}
