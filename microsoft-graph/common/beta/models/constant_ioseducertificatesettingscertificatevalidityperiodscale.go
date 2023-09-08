package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEduCertificateSettingsCertificateValidityPeriodScale string

const (
	IosEduCertificateSettingsCertificateValidityPeriodScaledays   IosEduCertificateSettingsCertificateValidityPeriodScale = "Days"
	IosEduCertificateSettingsCertificateValidityPeriodScalemonths IosEduCertificateSettingsCertificateValidityPeriodScale = "Months"
	IosEduCertificateSettingsCertificateValidityPeriodScaleyears  IosEduCertificateSettingsCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForIosEduCertificateSettingsCertificateValidityPeriodScale() []string {
	return []string{
		string(IosEduCertificateSettingsCertificateValidityPeriodScaledays),
		string(IosEduCertificateSettingsCertificateValidityPeriodScalemonths),
		string(IosEduCertificateSettingsCertificateValidityPeriodScaleyears),
	}
}

func parseIosEduCertificateSettingsCertificateValidityPeriodScale(input string) (*IosEduCertificateSettingsCertificateValidityPeriodScale, error) {
	vals := map[string]IosEduCertificateSettingsCertificateValidityPeriodScale{
		"days":   IosEduCertificateSettingsCertificateValidityPeriodScaledays,
		"months": IosEduCertificateSettingsCertificateValidityPeriodScalemonths,
		"years":  IosEduCertificateSettingsCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEduCertificateSettingsCertificateValidityPeriodScale(input)
	return &out, nil
}
