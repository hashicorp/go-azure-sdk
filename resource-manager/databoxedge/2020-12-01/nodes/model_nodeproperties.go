package nodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeProperties struct {
	NodeChassisSerialNumber     *string     `json:"nodeChassisSerialNumber,omitempty"`
	NodeDisplayName             *string     `json:"nodeDisplayName,omitempty"`
	NodeFriendlySoftwareVersion *string     `json:"nodeFriendlySoftwareVersion,omitempty"`
	NodeHcsVersion              *string     `json:"nodeHcsVersion,omitempty"`
	NodeInstanceId              *string     `json:"nodeInstanceId,omitempty"`
	NodeSerialNumber            *string     `json:"nodeSerialNumber,omitempty"`
	NodeStatus                  *NodeStatus `json:"nodeStatus,omitempty"`
}
