package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEduCertificateSettingsCertificateValidityPeriodScale string

const (
	IosEduCertificateSettingsCertificateValidityPeriodScale_Days   IosEduCertificateSettingsCertificateValidityPeriodScale = "days"
	IosEduCertificateSettingsCertificateValidityPeriodScale_Months IosEduCertificateSettingsCertificateValidityPeriodScale = "months"
	IosEduCertificateSettingsCertificateValidityPeriodScale_Years  IosEduCertificateSettingsCertificateValidityPeriodScale = "years"
)

func PossibleValuesForIosEduCertificateSettingsCertificateValidityPeriodScale() []string {
	return []string{
		string(IosEduCertificateSettingsCertificateValidityPeriodScale_Days),
		string(IosEduCertificateSettingsCertificateValidityPeriodScale_Months),
		string(IosEduCertificateSettingsCertificateValidityPeriodScale_Years),
	}
}

func (s *IosEduCertificateSettingsCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEduCertificateSettingsCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEduCertificateSettingsCertificateValidityPeriodScale(input string) (*IosEduCertificateSettingsCertificateValidityPeriodScale, error) {
	vals := map[string]IosEduCertificateSettingsCertificateValidityPeriodScale{
		"days":   IosEduCertificateSettingsCertificateValidityPeriodScale_Days,
		"months": IosEduCertificateSettingsCertificateValidityPeriodScale_Months,
		"years":  IosEduCertificateSettingsCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEduCertificateSettingsCertificateValidityPeriodScale(input)
	return &out, nil
}
