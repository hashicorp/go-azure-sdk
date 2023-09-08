package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseCertificateValidityPeriodScale string

const (
	MacOSCertificateProfileBaseCertificateValidityPeriodScaledays   MacOSCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	MacOSCertificateProfileBaseCertificateValidityPeriodScalemonths MacOSCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	MacOSCertificateProfileBaseCertificateValidityPeriodScaleyears  MacOSCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForMacOSCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseMacOSCertificateProfileBaseCertificateValidityPeriodScale(input string) (*MacOSCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   MacOSCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": MacOSCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  MacOSCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
