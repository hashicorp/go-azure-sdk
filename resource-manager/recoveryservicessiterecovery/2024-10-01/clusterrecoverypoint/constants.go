package clusterrecoverypoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterRecoveryPointType string

const (
	ClusterRecoveryPointTypeApplicationConsistent ClusterRecoveryPointType = "ApplicationConsistent"
	ClusterRecoveryPointTypeCrashConsistent       ClusterRecoveryPointType = "CrashConsistent"
	ClusterRecoveryPointTypeNotSpecified          ClusterRecoveryPointType = "NotSpecified"
)

func PossibleValuesForClusterRecoveryPointType() []string {
	return []string{
		string(ClusterRecoveryPointTypeApplicationConsistent),
		string(ClusterRecoveryPointTypeCrashConsistent),
		string(ClusterRecoveryPointTypeNotSpecified),
	}
}

func (s *ClusterRecoveryPointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterRecoveryPointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterRecoveryPointType(input string) (*ClusterRecoveryPointType, error) {
	vals := map[string]ClusterRecoveryPointType{
		"applicationconsistent": ClusterRecoveryPointTypeApplicationConsistent,
		"crashconsistent":       ClusterRecoveryPointTypeCrashConsistent,
		"notspecified":          ClusterRecoveryPointTypeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterRecoveryPointType(input)
	return &out, nil
}

type RecoveryPointSyncType string

const (
	RecoveryPointSyncTypeMultiVMSyncRecoveryPoint RecoveryPointSyncType = "MultiVmSyncRecoveryPoint"
)

func PossibleValuesForRecoveryPointSyncType() []string {
	return []string{
		string(RecoveryPointSyncTypeMultiVMSyncRecoveryPoint),
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
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointSyncType(input)
	return &out, nil
}
