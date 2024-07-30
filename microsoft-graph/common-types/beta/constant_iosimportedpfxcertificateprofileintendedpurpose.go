package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosImportedPFXCertificateProfileIntendedPurpose string

const (
	IosImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption IosImportedPFXCertificateProfileIntendedPurpose = "smimeEncryption"
	IosImportedPFXCertificateProfileIntendedPurpose_SmimeSigning    IosImportedPFXCertificateProfileIntendedPurpose = "smimeSigning"
	IosImportedPFXCertificateProfileIntendedPurpose_Unassigned      IosImportedPFXCertificateProfileIntendedPurpose = "unassigned"
	IosImportedPFXCertificateProfileIntendedPurpose_Vpn             IosImportedPFXCertificateProfileIntendedPurpose = "vpn"
	IosImportedPFXCertificateProfileIntendedPurpose_Wifi            IosImportedPFXCertificateProfileIntendedPurpose = "wifi"
)

func PossibleValuesForIosImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(IosImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption),
		string(IosImportedPFXCertificateProfileIntendedPurpose_SmimeSigning),
		string(IosImportedPFXCertificateProfileIntendedPurpose_Unassigned),
		string(IosImportedPFXCertificateProfileIntendedPurpose_Vpn),
		string(IosImportedPFXCertificateProfileIntendedPurpose_Wifi),
	}
}

func (s *IosImportedPFXCertificateProfileIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosImportedPFXCertificateProfileIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosImportedPFXCertificateProfileIntendedPurpose(input string) (*IosImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]IosImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": IosImportedPFXCertificateProfileIntendedPurpose_SmimeEncryption,
		"smimesigning":    IosImportedPFXCertificateProfileIntendedPurpose_SmimeSigning,
		"unassigned":      IosImportedPFXCertificateProfileIntendedPurpose_Unassigned,
		"vpn":             IosImportedPFXCertificateProfileIntendedPurpose_Vpn,
		"wifi":            IosImportedPFXCertificateProfileIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
