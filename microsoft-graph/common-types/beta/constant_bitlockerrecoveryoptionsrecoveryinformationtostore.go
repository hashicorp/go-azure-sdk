package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryInformationToStore string

const (
	BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordAndKey BitLockerRecoveryOptionsRecoveryInformationToStore = "passwordAndKey"
	BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordOnly   BitLockerRecoveryOptionsRecoveryInformationToStore = "passwordOnly"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryInformationToStore() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordAndKey),
		string(BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordOnly),
	}
}

func (s *BitLockerRecoveryOptionsRecoveryInformationToStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRecoveryOptionsRecoveryInformationToStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRecoveryOptionsRecoveryInformationToStore(input string) (*BitLockerRecoveryOptionsRecoveryInformationToStore, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryInformationToStore{
		"passwordandkey": BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordAndKey,
		"passwordonly":   BitLockerRecoveryOptionsRecoveryInformationToStore_PasswordOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryInformationToStore(input)
	return &out, nil
}
