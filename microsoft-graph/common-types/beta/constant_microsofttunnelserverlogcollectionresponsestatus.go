package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerLogCollectionResponseStatus string

const (
	MicrosoftTunnelServerLogCollectionResponseStatus_Completed MicrosoftTunnelServerLogCollectionResponseStatus = "completed"
	MicrosoftTunnelServerLogCollectionResponseStatus_Failed    MicrosoftTunnelServerLogCollectionResponseStatus = "failed"
	MicrosoftTunnelServerLogCollectionResponseStatus_Pending   MicrosoftTunnelServerLogCollectionResponseStatus = "pending"
)

func PossibleValuesForMicrosoftTunnelServerLogCollectionResponseStatus() []string {
	return []string{
		string(MicrosoftTunnelServerLogCollectionResponseStatus_Completed),
		string(MicrosoftTunnelServerLogCollectionResponseStatus_Failed),
		string(MicrosoftTunnelServerLogCollectionResponseStatus_Pending),
	}
}

func (s *MicrosoftTunnelServerLogCollectionResponseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTunnelServerLogCollectionResponseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTunnelServerLogCollectionResponseStatus(input string) (*MicrosoftTunnelServerLogCollectionResponseStatus, error) {
	vals := map[string]MicrosoftTunnelServerLogCollectionResponseStatus{
		"completed": MicrosoftTunnelServerLogCollectionResponseStatus_Completed,
		"failed":    MicrosoftTunnelServerLogCollectionResponseStatus_Failed,
		"pending":   MicrosoftTunnelServerLogCollectionResponseStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelServerLogCollectionResponseStatus(input)
	return &out, nil
}
