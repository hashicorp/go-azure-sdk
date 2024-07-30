package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Unassigned      AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Vpn             AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "vpn"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Wifi            AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
