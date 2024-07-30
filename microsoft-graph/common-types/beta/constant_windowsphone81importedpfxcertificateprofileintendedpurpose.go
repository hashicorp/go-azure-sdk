package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose string

const (
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Unassigned      WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Vpn             WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "vpn"
	WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Wifi            WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81ImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81ImportedPFXCertificateProfileIntendedPurpose(input string) (*WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
