package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Days   AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "days"
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Months AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "months"
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Years  AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
