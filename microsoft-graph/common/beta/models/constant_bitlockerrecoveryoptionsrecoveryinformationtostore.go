package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryInformationToStore string

const (
	BitLockerRecoveryOptionsRecoveryInformationToStorepasswordAndKey BitLockerRecoveryOptionsRecoveryInformationToStore = "PasswordAndKey"
	BitLockerRecoveryOptionsRecoveryInformationToStorepasswordOnly   BitLockerRecoveryOptionsRecoveryInformationToStore = "PasswordOnly"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryInformationToStore() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryInformationToStorepasswordAndKey),
		string(BitLockerRecoveryOptionsRecoveryInformationToStorepasswordOnly),
	}
}

func parseBitLockerRecoveryOptionsRecoveryInformationToStore(input string) (*BitLockerRecoveryOptionsRecoveryInformationToStore, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryInformationToStore{
		"passwordandkey": BitLockerRecoveryOptionsRecoveryInformationToStorepasswordAndKey,
		"passwordonly":   BitLockerRecoveryOptionsRecoveryInformationToStorepasswordOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryInformationToStore(input)
	return &out, nil
}
