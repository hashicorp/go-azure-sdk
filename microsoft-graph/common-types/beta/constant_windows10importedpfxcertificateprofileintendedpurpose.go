package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileIntendedPurpose string

const (
	Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption Windows10ImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    Windows10ImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	Windows10ImportedPFXCertificateProfileIntendedPurpose_Unassigned      Windows10ImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	Windows10ImportedPFXCertificateProfileIntendedPurpose_Vpn             Windows10ImportedPFXCertificateProfileIntendedPurpose = "vpn"
	Windows10ImportedPFXCertificateProfileIntendedPurpose_Wifi            Windows10ImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(Windows10ImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(Windows10ImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(Windows10ImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *Windows10ImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10ImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10ImportedPFXCertificateProfileIntendedPurpose(input string) (*Windows10ImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    Windows10ImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      Windows10ImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             Windows10ImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            Windows10ImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
