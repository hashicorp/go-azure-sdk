package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendDtmfTonesOperationCompletionReason string

const (
	SendDtmfTonesOperationCompletionReason_CompletedSuccessfully  SendDtmfTonesOperationCompletionReason = "completedSuccessfully"
	SendDtmfTonesOperationCompletionReason_MediaOperationCanceled SendDtmfTonesOperationCompletionReason = "mediaOperationCanceled"
	SendDtmfTonesOperationCompletionReason_Unknown                SendDtmfTonesOperationCompletionReason = "unknown"
)

func PossibleValuesForSendDtmfTonesOperationCompletionReason() []string {
	return []string{
		string(SendDtmfTonesOperationCompletionReason_CompletedSuccessfully),
		string(SendDtmfTonesOperationCompletionReason_MediaOperationCanceled),
		string(SendDtmfTonesOperationCompletionReason_Unknown),
	}
}

func (s *SendDtmfTonesOperationCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendDtmfTonesOperationCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendDtmfTonesOperationCompletionReason(input string) (*SendDtmfTonesOperationCompletionReason, error) {
	vals := map[string]SendDtmfTonesOperationCompletionReason{
		"completedsuccessfully":  SendDtmfTonesOperationCompletionReason_CompletedSuccessfully,
		"mediaoperationcanceled": SendDtmfTonesOperationCompletionReason_MediaOperationCanceled,
		"unknown":                SendDtmfTonesOperationCompletionReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendDtmfTonesOperationCompletionReason(input)
	return &out, nil
}
