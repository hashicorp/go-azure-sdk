package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileCertificateValidityPeriodScale string

const (
	Windows81SCEPCertificateProfileCertificateValidityPeriodScaledays   Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "Days"
	Windows81SCEPCertificateProfileCertificateValidityPeriodScalemonths Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "Months"
	Windows81SCEPCertificateProfileCertificateValidityPeriodScaleyears  Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows81SCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScaledays),
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScalemonths),
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseWindows81SCEPCertificateProfileCertificateValidityPeriodScale(input string) (*Windows81SCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows81SCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows81SCEPCertificateProfileCertificateValidityPeriodScaledays,
		"months": Windows81SCEPCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  Windows81SCEPCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
