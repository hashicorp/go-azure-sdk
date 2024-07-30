package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Allowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Blocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_NotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "notConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Required      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Allowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Blocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_NotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Required),
	}
}

func (s *BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Allowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Blocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_NotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage(input)
	return &out, nil
}
