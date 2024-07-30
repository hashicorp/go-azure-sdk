package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Days   AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "days"
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Months AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "months"
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Years  AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
