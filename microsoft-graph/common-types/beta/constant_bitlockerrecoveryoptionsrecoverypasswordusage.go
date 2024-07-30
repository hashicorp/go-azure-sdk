package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryPasswordUsage string

const (
	BitLockerRecoveryOptionsRecoveryPasswordUsage_Allowed       BitLockerRecoveryOptionsRecoveryPasswordUsage = "allowed"
	BitLockerRecoveryOptionsRecoveryPasswordUsage_Blocked       BitLockerRecoveryOptionsRecoveryPasswordUsage = "blocked"
	BitLockerRecoveryOptionsRecoveryPasswordUsage_NotConfigured BitLockerRecoveryOptionsRecoveryPasswordUsage = "notConfigured"
	BitLockerRecoveryOptionsRecoveryPasswordUsage_Required      BitLockerRecoveryOptionsRecoveryPasswordUsage = "required"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryPasswordUsage() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryPasswordUsage_Allowed),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsage_Blocked),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsage_NotConfigured),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsage_Required),
	}
}

func (s *BitLockerRecoveryOptionsRecoveryPasswordUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRecoveryOptionsRecoveryPasswordUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRecoveryOptionsRecoveryPasswordUsage(input string) (*BitLockerRecoveryOptionsRecoveryPasswordUsage, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryPasswordUsage{
		"allowed":       BitLockerRecoveryOptionsRecoveryPasswordUsage_Allowed,
		"blocked":       BitLockerRecoveryOptionsRecoveryPasswordUsage_Blocked,
		"notconfigured": BitLockerRecoveryOptionsRecoveryPasswordUsage_NotConfigured,
		"required":      BitLockerRecoveryOptionsRecoveryPasswordUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryPasswordUsage(input)
	return &out, nil
}
