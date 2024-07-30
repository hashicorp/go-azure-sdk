package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Unassigned      AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Vpn             AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "vpn"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Wifi            AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *AndroidForWorkImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidForWorkImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            AndroidForWorkImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
