package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Days   AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Months AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Years  AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
