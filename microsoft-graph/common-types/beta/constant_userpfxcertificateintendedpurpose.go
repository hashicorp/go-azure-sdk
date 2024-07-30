package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificateIntendedPurpose string

const (
	UserPFXCertificateIntendedPurpose_SmimeEncryption UserPFXCertificateIntendedPurpose = "smimeEncryption"
	UserPFXCertificateIntendedPurpose_SmimeSigning    UserPFXCertificateIntendedPurpose = "smimeSigning"
	UserPFXCertificateIntendedPurpose_Unassigned      UserPFXCertificateIntendedPurpose = "unassigned"
	UserPFXCertificateIntendedPurpose_Vpn             UserPFXCertificateIntendedPurpose = "vpn"
	UserPFXCertificateIntendedPurpose_Wifi            UserPFXCertificateIntendedPurpose = "wifi"
)

func PossibleValuesForUserPFXCertificateIntendedPurpose() []string {
	return []string{
		string(UserPFXCertificateIntendedPurpose_SmimeEncryption),
		string(UserPFXCertificateIntendedPurpose_SmimeSigning),
		string(UserPFXCertificateIntendedPurpose_Unassigned),
		string(UserPFXCertificateIntendedPurpose_Vpn),
		string(UserPFXCertificateIntendedPurpose_Wifi),
	}
}

func (s *UserPFXCertificateIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserPFXCertificateIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserPFXCertificateIntendedPurpose(input string) (*UserPFXCertificateIntendedPurpose, error) {
	vals := map[string]UserPFXCertificateIntendedPurpose{
		"smimeencryption": UserPFXCertificateIntendedPurpose_SmimeEncryption,
		"smimesigning":    UserPFXCertificateIntendedPurpose_SmimeSigning,
		"unassigned":      UserPFXCertificateIntendedPurpose_Unassigned,
		"vpn":             UserPFXCertificateIntendedPurpose_Vpn,
		"wifi":            UserPFXCertificateIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPFXCertificateIntendedPurpose(input)
	return &out, nil
}
