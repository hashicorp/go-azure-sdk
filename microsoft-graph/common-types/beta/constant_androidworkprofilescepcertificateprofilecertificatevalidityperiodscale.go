package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Days   AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Months AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Years  AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
