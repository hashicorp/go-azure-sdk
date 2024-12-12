package sessionhostmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolUpdateControlParameter struct {
	Action        HostPoolUpdateAction `json:"action"`
	CancelMessage *string              `json:"cancelMessage,omitempty"`
}
