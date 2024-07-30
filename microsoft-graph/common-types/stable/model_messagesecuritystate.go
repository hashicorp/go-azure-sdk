package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageSecurityState struct {
	ConnectingIP            *string `json:"connectingIP,omitempty"`
	DeliveryAction          *string `json:"deliveryAction,omitempty"`
	DeliveryLocation        *string `json:"deliveryLocation,omitempty"`
	Directionality          *string `json:"directionality,omitempty"`
	InternetMessageId       *string `json:"internetMessageId,omitempty"`
	MessageFingerprint      *string `json:"messageFingerprint,omitempty"`
	MessageReceivedDateTime *string `json:"messageReceivedDateTime,omitempty"`
	MessageSubject          *string `json:"messageSubject,omitempty"`
	NetworkMessageId        *string `json:"networkMessageId,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
}
