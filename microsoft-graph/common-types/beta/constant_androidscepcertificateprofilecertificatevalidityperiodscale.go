package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidScepCertificateProfileCertificateValidityPeriodScale_Days   AndroidScepCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidScepCertificateProfileCertificateValidityPeriodScale_Months AndroidScepCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidScepCertificateProfileCertificateValidityPeriodScale_Years  AndroidScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
