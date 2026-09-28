package fleets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status string

const (
	StatusCanceled        Status = "Canceled"
	StatusCreating        Status = "Creating"
	StatusDeleting        Status = "Deleting"
	StatusFailed          Status = "Failed"
	StatusInitializing    Status = "Initializing"
	StatusInternallyReady Status = "InternallyReady"
	StatusOnline          Status = "Online"
	StatusSucceeded       Status = "Succeeded"
	StatusUninitialized   Status = "Uninitialized"
	StatusUpdating        Status = "Updating"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusCanceled),
		string(StatusCreating),
		string(StatusDeleting),
		string(StatusFailed),
		string(StatusInitializing),
		string(StatusInternallyReady),
		string(StatusOnline),
		string(StatusSucceeded),
		string(StatusUninitialized),
		string(StatusUpdating),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"canceled":        StatusCanceled,
		"creating":        StatusCreating,
		"deleting":        StatusDeleting,
		"failed":          StatusFailed,
		"initializing":    StatusInitializing,
		"internallyready": StatusInternallyReady,
		"online":          StatusOnline,
		"succeeded":       StatusSucceeded,
		"uninitialized":   StatusUninitialized,
		"updating":        StatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
