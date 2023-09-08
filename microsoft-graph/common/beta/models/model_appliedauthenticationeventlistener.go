package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedAuthenticationEventListener struct {
	EventType          *AppliedAuthenticationEventListenerEventType `json:"eventType,omitempty"`
	ExecutedListenerId *string                                      `json:"executedListenerId,omitempty"`
	HandlerResult      *AuthenticationEventHandlerResult            `json:"handlerResult,omitempty"`
	ODataType          *string                                      `json:"@odata.type,omitempty"`
}
