package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseCertificateValidityPeriodScale string

const (
	IosCertificateProfileBaseCertificateValidityPeriodScaledays   IosCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	IosCertificateProfileBaseCertificateValidityPeriodScalemonths IosCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	IosCertificateProfileBaseCertificateValidityPeriodScaleyears  IosCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForIosCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(IosCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(IosCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(IosCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseIosCertificateProfileBaseCertificateValidityPeriodScale(input string) (*IosCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]IosCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   IosCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": IosCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  IosCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
