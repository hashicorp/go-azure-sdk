package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	MacOSPkcsCertificateProfileCertificateValidityPeriodScaledays   MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	MacOSPkcsCertificateProfileCertificateValidityPeriodScalemonths MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	MacOSPkcsCertificateProfileCertificateValidityPeriodScaleyears  MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForMacOSPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseMacOSPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*MacOSPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   MacOSPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": MacOSPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  MacOSPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
