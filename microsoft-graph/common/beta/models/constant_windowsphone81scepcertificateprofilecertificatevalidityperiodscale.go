package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale string

const (
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaledays   WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "Days"
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScalemonths WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "Months"
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaleyears  WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaledays),
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScalemonths),
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseWindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale(input string) (*WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaledays,
		"months": WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
