package syncagents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncAgentState string

const (
	SyncAgentStateNeverConnected SyncAgentState = "NeverConnected"
	SyncAgentStateOffline        SyncAgentState = "Offline"
	SyncAgentStateOnline         SyncAgentState = "Online"
)

func PossibleValuesForSyncAgentState() []string {
	return []string{
		string(SyncAgentStateNeverConnected),
		string(SyncAgentStateOffline),
		string(SyncAgentStateOnline),
	}
}

func (s *SyncAgentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncAgentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncAgentState(input string) (*SyncAgentState, error) {
	vals := map[string]SyncAgentState{
		"neverconnected": SyncAgentStateNeverConnected,
		"offline":        SyncAgentStateOffline,
		"online":         SyncAgentStateOnline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncAgentState(input)
	return &out, nil
}

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  SyncMemberDbType = "AzureSqlDatabase"
	SyncMemberDbTypeSqlServerDatabase SyncMemberDbType = "SqlServerDatabase"
)

func PossibleValuesForSyncMemberDbType() []string {
	return []string{
		string(SyncMemberDbTypeAzureSqlDatabase),
		string(SyncMemberDbTypeSqlServerDatabase),
	}
}

func (s *SyncMemberDbType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncMemberDbType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncMemberDbType(input string) (*SyncMemberDbType, error) {
	vals := map[string]SyncMemberDbType{
		"azuresqldatabase":  SyncMemberDbTypeAzureSqlDatabase,
		"sqlserverdatabase": SyncMemberDbTypeSqlServerDatabase,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncMemberDbType(input)
	return &out, nil
}
