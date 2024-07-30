package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileCertificateValidityPeriodScale string

const (
	IosScepCertificateProfileCertificateValidityPeriodScale_Days   IosScepCertificateProfileCertificateValidityPeriodScale = "days"
	IosScepCertificateProfileCertificateValidityPeriodScale_Months IosScepCertificateProfileCertificateValidityPeriodScale = "months"
	IosScepCertificateProfileCertificateValidityPeriodScale_Years  IosScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForIosScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(IosScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(IosScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(IosScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *IosScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileCertificateValidityPeriodScale(input string) (*IosScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]IosScepCertificateProfileCertificateValidityPeriodScale{
		"days":   IosScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": IosScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  IosScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
