package watchlists

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceType string

const (
	SourceTypeAzureStorage SourceType = "AzureStorage"
	SourceTypeLocal        SourceType = "Local"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeAzureStorage),
		string(SourceTypeLocal),
	}
}

func (s *SourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSourceType(input string) (*SourceType, error) {
	vals := map[string]SourceType{
		"azurestorage": SourceTypeAzureStorage,
		"local":        SourceTypeLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceType(input)
	return &out, nil
}

type WatchlistProvisioningState string

const (
	WatchlistProvisioningStateCanceled   WatchlistProvisioningState = "Canceled"
	WatchlistProvisioningStateDeleting   WatchlistProvisioningState = "Deleting"
	WatchlistProvisioningStateFailed     WatchlistProvisioningState = "Failed"
	WatchlistProvisioningStateInProgress WatchlistProvisioningState = "InProgress"
	WatchlistProvisioningStateNew        WatchlistProvisioningState = "New"
	WatchlistProvisioningStateSucceeded  WatchlistProvisioningState = "Succeeded"
	WatchlistProvisioningStateUploading  WatchlistProvisioningState = "Uploading"
)

func PossibleValuesForWatchlistProvisioningState() []string {
	return []string{
		string(WatchlistProvisioningStateCanceled),
		string(WatchlistProvisioningStateDeleting),
		string(WatchlistProvisioningStateFailed),
		string(WatchlistProvisioningStateInProgress),
		string(WatchlistProvisioningStateNew),
		string(WatchlistProvisioningStateSucceeded),
		string(WatchlistProvisioningStateUploading),
	}
}

func (s *WatchlistProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWatchlistProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWatchlistProvisioningState(input string) (*WatchlistProvisioningState, error) {
	vals := map[string]WatchlistProvisioningState{
		"canceled":   WatchlistProvisioningStateCanceled,
		"deleting":   WatchlistProvisioningStateDeleting,
		"failed":     WatchlistProvisioningStateFailed,
		"inprogress": WatchlistProvisioningStateInProgress,
		"new":        WatchlistProvisioningStateNew,
		"succeeded":  WatchlistProvisioningStateSucceeded,
		"uploading":  WatchlistProvisioningStateUploading,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WatchlistProvisioningState(input)
	return &out, nil
}
