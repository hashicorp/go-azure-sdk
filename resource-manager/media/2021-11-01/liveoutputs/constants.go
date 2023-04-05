package liveoutputs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiveOutputResourceState string

const (
	LiveOutputResourceStateCreating LiveOutputResourceState = "Creating"
	LiveOutputResourceStateDeleting LiveOutputResourceState = "Deleting"
	LiveOutputResourceStateRunning  LiveOutputResourceState = "Running"
)

func PossibleValuesForLiveOutputResourceState() []string {
	return []string{
		string(LiveOutputResourceStateCreating),
		string(LiveOutputResourceStateDeleting),
		string(LiveOutputResourceStateRunning),
	}
}

func (s *LiveOutputResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLiveOutputResourceState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = LiveOutputResourceState(decoded)
	return nil
}
