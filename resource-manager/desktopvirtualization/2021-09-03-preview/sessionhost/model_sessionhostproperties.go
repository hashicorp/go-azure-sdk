package sessionhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostProperties struct {
	AgentVersion                  *string                         `json:"agentVersion,omitempty"`
	AllowNewSession               *bool                           `json:"allowNewSession,omitempty"`
	AssignedUser                  *string                         `json:"assignedUser,omitempty"`
	LastHeartBeat                 *string                         `json:"lastHeartBeat,omitempty"`
	LastUpdateTime                *string                         `json:"lastUpdateTime,omitempty"`
	ObjectId                      *string                         `json:"objectId,omitempty"`
	OsVersion                     *string                         `json:"osVersion,omitempty"`
	ResourceId                    *string                         `json:"resourceId,omitempty"`
	SessionHostHealthCheckResults *[]SessionHostHealthCheckReport `json:"sessionHostHealthCheckResults,omitempty"`
	Sessions                      *int64                          `json:"sessions,omitempty"`
	Status                        *Status                         `json:"status,omitempty"`
	StatusTimestamp               *string                         `json:"statusTimestamp,omitempty"`
	SxSStackVersion               *string                         `json:"sxSStackVersion,omitempty"`
	UpdateErrorMessage            *string                         `json:"updateErrorMessage,omitempty"`
	UpdateState                   *UpdateState                    `json:"updateState,omitempty"`
	VirtualMachineId              *string                         `json:"virtualMachineId,omitempty"`
}
