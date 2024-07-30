package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days   AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "days"
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "months"
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years  AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
