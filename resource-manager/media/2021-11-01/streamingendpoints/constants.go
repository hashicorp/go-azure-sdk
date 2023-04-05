package streamingendpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingEndpointResourceState string

const (
	StreamingEndpointResourceStateDeleting StreamingEndpointResourceState = "Deleting"
	StreamingEndpointResourceStateRunning  StreamingEndpointResourceState = "Running"
	StreamingEndpointResourceStateScaling  StreamingEndpointResourceState = "Scaling"
	StreamingEndpointResourceStateStarting StreamingEndpointResourceState = "Starting"
	StreamingEndpointResourceStateStopped  StreamingEndpointResourceState = "Stopped"
	StreamingEndpointResourceStateStopping StreamingEndpointResourceState = "Stopping"
)

func PossibleValuesForStreamingEndpointResourceState() []string {
	return []string{
		string(StreamingEndpointResourceStateDeleting),
		string(StreamingEndpointResourceStateRunning),
		string(StreamingEndpointResourceStateScaling),
		string(StreamingEndpointResourceStateStarting),
		string(StreamingEndpointResourceStateStopped),
		string(StreamingEndpointResourceStateStopping),
	}
}

func (s *StreamingEndpointResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStreamingEndpointResourceState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StreamingEndpointResourceState(decoded)
	return nil
}
