package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Allowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Blocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_NotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "notConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Required      BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Allowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Blocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_NotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Required),
	}
}

func (s *BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Allowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Blocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_NotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage(input)
	return &out, nil
}
