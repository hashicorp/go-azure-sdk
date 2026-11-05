package agentsconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorProvisioningState string

const (
	ConnectorProvisioningStateCanceled   ConnectorProvisioningState = "Canceled"
	ConnectorProvisioningStateDeleting   ConnectorProvisioningState = "Deleting"
	ConnectorProvisioningStateFailed     ConnectorProvisioningState = "Failed"
	ConnectorProvisioningStateInProgress ConnectorProvisioningState = "InProgress"
	ConnectorProvisioningStateSucceeded  ConnectorProvisioningState = "Succeeded"
)

func PossibleValuesForConnectorProvisioningState() []string {
	return []string{
		string(ConnectorProvisioningStateCanceled),
		string(ConnectorProvisioningStateDeleting),
		string(ConnectorProvisioningStateFailed),
		string(ConnectorProvisioningStateInProgress),
		string(ConnectorProvisioningStateSucceeded),
	}
}

func (s *ConnectorProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorProvisioningState(input string) (*ConnectorProvisioningState, error) {
	vals := map[string]ConnectorProvisioningState{
		"canceled":   ConnectorProvisioningStateCanceled,
		"deleting":   ConnectorProvisioningStateDeleting,
		"failed":     ConnectorProvisioningStateFailed,
		"inprogress": ConnectorProvisioningStateInProgress,
		"succeeded":  ConnectorProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorProvisioningState(input)
	return &out, nil
}
