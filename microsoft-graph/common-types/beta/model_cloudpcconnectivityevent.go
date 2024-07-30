package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEvent struct {
	EventDateTime *string                              `json:"eventDateTime,omitempty"`
	EventName     *string                              `json:"eventName,omitempty"`
	EventResult   *CloudPCConnectivityEventEventResult `json:"eventResult,omitempty"`
	EventType     *CloudPCConnectivityEventEventType   `json:"eventType,omitempty"`
	Message       *string                              `json:"message,omitempty"`
	ODataType     *string                              `json:"@odata.type,omitempty"`
}
