package recoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointSyncType string

const (
	RecoveryPointSyncTypeMultiVMSyncRecoveryPoint RecoveryPointSyncType = "MultiVmSyncRecoveryPoint"
	RecoveryPointSyncTypePerVMRecoveryPoint       RecoveryPointSyncType = "PerVmRecoveryPoint"
)

func PossibleValuesForRecoveryPointSyncType() []string {
	return []string{
		string(RecoveryPointSyncTypeMultiVMSyncRecoveryPoint),
		string(RecoveryPointSyncTypePerVMRecoveryPoint),
	}
}

func (s *RecoveryPointSyncType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecoveryPointSyncType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecoveryPointSyncType(input string) (*RecoveryPointSyncType, error) {
	vals := map[string]RecoveryPointSyncType{
		"multivmsyncrecoverypoint": RecoveryPointSyncTypeMultiVMSyncRecoveryPoint,
		"pervmrecoverypoint":       RecoveryPointSyncTypePerVMRecoveryPoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointSyncType(input)
	return &out, nil
}
