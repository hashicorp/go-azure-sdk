package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVppTokenTroubleshootingEvent struct {
	AdditionalInformation       *[]KeyValuePair                              `json:"additionalInformation,omitempty"`
	CorrelationId               *string                                      `json:"correlationId,omitempty"`
	EventDateTime               *string                                      `json:"eventDateTime,omitempty"`
	EventName                   *string                                      `json:"eventName,omitempty"`
	Id                          *string                                      `json:"id,omitempty"`
	ODataType                   *string                                      `json:"@odata.type,omitempty"`
	TokenId                     *string                                      `json:"tokenId,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`
}
