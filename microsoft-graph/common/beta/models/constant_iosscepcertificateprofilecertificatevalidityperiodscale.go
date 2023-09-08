package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileCertificateValidityPeriodScale string

const (
	IosScepCertificateProfileCertificateValidityPeriodScaledays   IosScepCertificateProfileCertificateValidityPeriodScale = "Days"
	IosScepCertificateProfileCertificateValidityPeriodScalemonths IosScepCertificateProfileCertificateValidityPeriodScale = "Months"
	IosScepCertificateProfileCertificateValidityPeriodScaleyears  IosScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForIosScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(IosScepCertificateProfileCertificateValidityPeriodScaledays),
		string(IosScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(IosScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseIosScepCertificateProfileCertificateValidityPeriodScale(input string) (*IosScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]IosScepCertificateProfileCertificateValidityPeriodScale{
		"days":   IosScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": IosScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  IosScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
