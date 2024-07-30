package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseCertificateValidityPeriodScale string

const (
	WindowsCertificateProfileBaseCertificateValidityPeriodScale_Days   WindowsCertificateProfileBaseCertificateValidityPeriodScale = "days"
	WindowsCertificateProfileBaseCertificateValidityPeriodScale_Months WindowsCertificateProfileBaseCertificateValidityPeriodScale = "months"
	WindowsCertificateProfileBaseCertificateValidityPeriodScale_Years  WindowsCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindowsCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *WindowsCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsCertificateProfileBaseCertificateValidityPeriodScale(input string) (*WindowsCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   WindowsCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": WindowsCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  WindowsCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
