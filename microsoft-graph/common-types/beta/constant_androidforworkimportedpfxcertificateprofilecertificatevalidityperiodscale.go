package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Months AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
