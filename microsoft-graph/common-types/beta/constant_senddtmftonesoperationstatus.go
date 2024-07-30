package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendDtmfTonesOperationStatus string

const (
	SendDtmfTonesOperationStatus_Completed  SendDtmfTonesOperationStatus = "Completed"
	SendDtmfTonesOperationStatus_Failed     SendDtmfTonesOperationStatus = "Failed"
	SendDtmfTonesOperationStatus_NotStarted SendDtmfTonesOperationStatus = "NotStarted"
	SendDtmfTonesOperationStatus_Running    SendDtmfTonesOperationStatus = "Running"
)

func PossibleValuesForSendDtmfTonesOperationStatus() []string {
	return []string{
		string(SendDtmfTonesOperationStatus_Completed),
		string(SendDtmfTonesOperationStatus_Failed),
		string(SendDtmfTonesOperationStatus_NotStarted),
		string(SendDtmfTonesOperationStatus_Running),
	}
}

func (s *SendDtmfTonesOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendDtmfTonesOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendDtmfTonesOperationStatus(input string) (*SendDtmfTonesOperationStatus, error) {
	vals := map[string]SendDtmfTonesOperationStatus{
		"completed":  SendDtmfTonesOperationStatus_Completed,
		"failed":     SendDtmfTonesOperationStatus_Failed,
		"notstarted": SendDtmfTonesOperationStatus_NotStarted,
		"running":    SendDtmfTonesOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendDtmfTonesOperationStatus(input)
	return &out, nil
}
