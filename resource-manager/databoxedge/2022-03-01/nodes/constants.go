package nodes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeStatus string

const (
	NodeStatusDown         NodeStatus = "Down"
	NodeStatusRebooting    NodeStatus = "Rebooting"
	NodeStatusShuttingDown NodeStatus = "ShuttingDown"
	NodeStatusUnknown      NodeStatus = "Unknown"
	NodeStatusUp           NodeStatus = "Up"
)

func PossibleValuesForNodeStatus() []string {
	return []string{
		string(NodeStatusDown),
		string(NodeStatusRebooting),
		string(NodeStatusShuttingDown),
		string(NodeStatusUnknown),
		string(NodeStatusUp),
	}
}

func (s *NodeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNodeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNodeStatus(input string) (*NodeStatus, error) {
	vals := map[string]NodeStatus{
		"down":         NodeStatusDown,
		"rebooting":    NodeStatusRebooting,
		"shuttingdown": NodeStatusShuttingDown,
		"unknown":      NodeStatusUnknown,
		"up":           NodeStatusUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeStatus(input)
	return &out, nil
}
