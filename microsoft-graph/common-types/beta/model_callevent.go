package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEvent struct {
	CallEventType *CallEventCallEventType `json:"callEventType,omitempty"`
	Direction     *CallEventDirection     `json:"direction,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	JoinCallUrl   *string                 `json:"joinCallUrl,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
}
