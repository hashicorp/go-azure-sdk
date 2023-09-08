package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceSettings struct {
	AllowSessionsToUnenrolledDevices *bool                                          `json:"allowSessionsToUnenrolledDevices,omitempty"`
	BlockChat                        *bool                                          `json:"blockChat,omitempty"`
	Id                               *string                                        `json:"id,omitempty"`
	ODataType                        *string                                        `json:"@odata.type,omitempty"`
	RemoteAssistanceState            *RemoteAssistanceSettingsRemoteAssistanceState `json:"remoteAssistanceState,omitempty"`
}
