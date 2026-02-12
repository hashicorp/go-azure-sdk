package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCommand struct {
	ConnectionTimeout *int64       `json:"connectionTimeout,omitempty"`
	Id                *string      `json:"id,omitempty"`
	Request           *interface{} `json:"request,omitempty"`
	Response          *interface{} `json:"response,omitempty"`
	ResponseCode      *int64       `json:"responseCode,omitempty"`
	ResponseTimeout   *int64       `json:"responseTimeout,omitempty"`
}
