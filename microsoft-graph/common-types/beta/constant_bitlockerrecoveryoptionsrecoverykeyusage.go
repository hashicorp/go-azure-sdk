package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryKeyUsage string

const (
	BitLockerRecoveryOptionsRecoveryKeyUsage_Allowed       BitLockerRecoveryOptionsRecoveryKeyUsage = "allowed"
	BitLockerRecoveryOptionsRecoveryKeyUsage_Blocked       BitLockerRecoveryOptionsRecoveryKeyUsage = "blocked"
	BitLockerRecoveryOptionsRecoveryKeyUsage_NotConfigured BitLockerRecoveryOptionsRecoveryKeyUsage = "notConfigured"
	BitLockerRecoveryOptionsRecoveryKeyUsage_Required      BitLockerRecoveryOptionsRecoveryKeyUsage = "required"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryKeyUsage() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryKeyUsage_Allowed),
		string(BitLockerRecoveryOptionsRecoveryKeyUsage_Blocked),
		string(BitLockerRecoveryOptionsRecoveryKeyUsage_NotConfigured),
		string(BitLockerRecoveryOptionsRecoveryKeyUsage_Required),
	}
}

func (s *BitLockerRecoveryOptionsRecoveryKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRecoveryOptionsRecoveryKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRecoveryOptionsRecoveryKeyUsage(input string) (*BitLockerRecoveryOptionsRecoveryKeyUsage, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryKeyUsage{
		"allowed":       BitLockerRecoveryOptionsRecoveryKeyUsage_Allowed,
		"blocked":       BitLockerRecoveryOptionsRecoveryKeyUsage_Blocked,
		"notconfigured": BitLockerRecoveryOptionsRecoveryKeyUsage_NotConfigured,
		"required":      BitLockerRecoveryOptionsRecoveryKeyUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryKeyUsage(input)
	return &out, nil
}
