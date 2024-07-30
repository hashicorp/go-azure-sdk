package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileIntendedPurpose string

const (
	MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption MacOSImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    MacOSImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	MacOSImportedPFXCertificateProfileIntendedPurpose_Unassigned      MacOSImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	MacOSImportedPFXCertificateProfileIntendedPurpose_Vpn             MacOSImportedPFXCertificateProfileIntendedPurpose = "vpn"
	MacOSImportedPFXCertificateProfileIntendedPurpose_Wifi            MacOSImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(MacOSImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(MacOSImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(MacOSImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *MacOSImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSImportedPFXCertificateProfileIntendedPurpose(input string) (*MacOSImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    MacOSImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      MacOSImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             MacOSImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            MacOSImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
