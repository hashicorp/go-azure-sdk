package sourcecontrolsyncjobstreams

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamType string

const (
	StreamTypeError  StreamType = "Error"
	StreamTypeOutput StreamType = "Output"
)

func PossibleValuesForStreamType() []string {
	return []string{
		string(StreamTypeError),
		string(StreamTypeOutput),
	}
}

func (s *StreamType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStreamType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StreamType(decoded)
	return nil
}
