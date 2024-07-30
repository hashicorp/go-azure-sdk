package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseCertificateValidityPeriodScale string

const (
	IosCertificateProfileBaseCertificateValidityPeriodScale_Days   IosCertificateProfileBaseCertificateValidityPeriodScale = "days"
	IosCertificateProfileBaseCertificateValidityPeriodScale_Months IosCertificateProfileBaseCertificateValidityPeriodScale = "months"
	IosCertificateProfileBaseCertificateValidityPeriodScale_Years  IosCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForIosCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(IosCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(IosCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(IosCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *IosCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCertificateProfileBaseCertificateValidityPeriodScale(input string) (*IosCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]IosCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   IosCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": IosCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  IosCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
