package monitoredresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendingLogs string

const (
	SendingLogsFalse SendingLogs = "False"
	SendingLogsTrue  SendingLogs = "True"
)

func PossibleValuesForSendingLogs() []string {
	return []string{
		string(SendingLogsFalse),
		string(SendingLogsTrue),
	}
}

func (s *SendingLogs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendingLogs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendingLogs(input string) (*SendingLogs, error) {
	vals := map[string]SendingLogs{
		"false": SendingLogsFalse,
		"true":  SendingLogsTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendingLogs(input)
	return &out, nil
}
