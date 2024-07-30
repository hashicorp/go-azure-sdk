package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption AndroidImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    AndroidImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	AndroidImportedPFXCertificateProfileIntendedPurpose_Unassigned      AndroidImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	AndroidImportedPFXCertificateProfileIntendedPurpose_Vpn             AndroidImportedPFXCertificateProfileIntendedPurpose = "vpn"
	AndroidImportedPFXCertificateProfileIntendedPurpose_Wifi            AndroidImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(AndroidImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(AndroidImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(AndroidImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *AndroidImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    AndroidImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      AndroidImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             AndroidImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            AndroidImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
