package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	IosPkcsCertificateProfileCertificateValidityPeriodScaledays   IosPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	IosPkcsCertificateProfileCertificateValidityPeriodScalemonths IosPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	IosPkcsCertificateProfileCertificateValidityPeriodScaleyears  IosPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForIosPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(IosPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(IosPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(IosPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseIosPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*IosPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]IosPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   IosPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": IosPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  IosPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
