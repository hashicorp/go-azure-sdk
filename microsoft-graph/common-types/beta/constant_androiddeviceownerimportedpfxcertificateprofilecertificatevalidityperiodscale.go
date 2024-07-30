package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Months AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
