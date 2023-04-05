package nodes

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
