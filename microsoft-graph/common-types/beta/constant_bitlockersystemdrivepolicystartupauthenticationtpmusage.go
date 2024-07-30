package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Allowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage = "allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Blocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage = "blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_NotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage = "notConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Required      BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage = "required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Allowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Blocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_NotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Required),
	}
}

func (s *BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerSystemDrivePolicyStartupAuthenticationTpmUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Allowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Blocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_NotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage(input)
	return &out, nil
}
