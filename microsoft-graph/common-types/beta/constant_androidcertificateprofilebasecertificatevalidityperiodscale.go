package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidCertificateProfileBaseCertificateValidityPeriodScale_Days   AndroidCertificateProfileBaseCertificateValidityPeriodScale = "days"
	AndroidCertificateProfileBaseCertificateValidityPeriodScale_Months AndroidCertificateProfileBaseCertificateValidityPeriodScale = "months"
	AndroidCertificateProfileBaseCertificateValidityPeriodScale_Years  AndroidCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": AndroidCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  AndroidCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
