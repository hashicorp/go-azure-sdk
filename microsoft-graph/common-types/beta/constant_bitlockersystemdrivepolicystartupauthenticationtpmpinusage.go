package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Allowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Blocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_NotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "notConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Required      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Allowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Blocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_NotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Required),
	}
}

func (s *BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Allowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Blocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_NotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage(input)
	return &out, nil
}
