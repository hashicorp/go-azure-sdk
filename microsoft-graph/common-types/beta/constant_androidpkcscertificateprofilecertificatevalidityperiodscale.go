package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Days   AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Months AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Years  AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
