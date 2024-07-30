package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Days   AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Months AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Years  AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
