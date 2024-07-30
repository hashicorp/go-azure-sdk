package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadStatus string

const (
	PayloadStatus_Archive PayloadStatus = "archive"
	PayloadStatus_Delete  PayloadStatus = "delete"
	PayloadStatus_Draft   PayloadStatus = "draft"
	PayloadStatus_Ready   PayloadStatus = "ready"
	PayloadStatus_Unknown PayloadStatus = "unknown"
)

func PossibleValuesForPayloadStatus() []string {
	return []string{
		string(PayloadStatus_Archive),
		string(PayloadStatus_Delete),
		string(PayloadStatus_Draft),
		string(PayloadStatus_Ready),
		string(PayloadStatus_Unknown),
	}
}

func (s *PayloadStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadStatus(input string) (*PayloadStatus, error) {
	vals := map[string]PayloadStatus{
		"archive": PayloadStatus_Archive,
		"delete":  PayloadStatus_Delete,
		"draft":   PayloadStatus_Draft,
		"ready":   PayloadStatus_Ready,
		"unknown": PayloadStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadStatus(input)
	return &out, nil
}
