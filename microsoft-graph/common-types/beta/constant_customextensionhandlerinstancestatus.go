package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstanceStatus string

const (
	CustomExtensionHandlerInstanceStatus_RequestReceived CustomExtensionHandlerInstanceStatus = "requestReceived"
	CustomExtensionHandlerInstanceStatus_RequestSent     CustomExtensionHandlerInstanceStatus = "requestSent"
)

func PossibleValuesForCustomExtensionHandlerInstanceStatus() []string {
	return []string{
		string(CustomExtensionHandlerInstanceStatus_RequestReceived),
		string(CustomExtensionHandlerInstanceStatus_RequestSent),
	}
}

func (s *CustomExtensionHandlerInstanceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomExtensionHandlerInstanceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomExtensionHandlerInstanceStatus(input string) (*CustomExtensionHandlerInstanceStatus, error) {
	vals := map[string]CustomExtensionHandlerInstanceStatus{
		"requestreceived": CustomExtensionHandlerInstanceStatus_RequestReceived,
		"requestsent":     CustomExtensionHandlerInstanceStatus_RequestSent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerInstanceStatus(input)
	return &out, nil
}
