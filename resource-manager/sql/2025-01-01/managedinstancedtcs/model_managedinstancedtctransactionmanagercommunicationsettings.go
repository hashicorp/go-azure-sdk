package managedinstancedtcs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceDtcTransactionManagerCommunicationSettings struct {
	AllowInboundEnabled  *bool   `json:"allowInboundEnabled,omitempty"`
	AllowOutboundEnabled *bool   `json:"allowOutboundEnabled,omitempty"`
	Authentication       *string `json:"authentication,omitempty"`
}
