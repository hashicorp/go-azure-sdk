package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingEvent struct {
	AdditionalInformation       *[]KeyValuePair                              `json:"additionalInformation,omitempty"`
	AppLogCollectionRequests    *[]AppLogCollectionRequest                   `json:"appLogCollectionRequests,omitempty"`
	ApplicationId               *string                                      `json:"applicationId,omitempty"`
	CorrelationId               *string                                      `json:"correlationId,omitempty"`
	DeviceId                    *string                                      `json:"deviceId,omitempty"`
	EventDateTime               *string                                      `json:"eventDateTime,omitempty"`
	EventName                   *string                                      `json:"eventName,omitempty"`
	History                     *[]MobileAppTroubleshootingHistoryItem       `json:"history,omitempty"`
	Id                          *string                                      `json:"id,omitempty"`
	ManagedDeviceIdentifier     *string                                      `json:"managedDeviceIdentifier,omitempty"`
	ODataType                   *string                                      `json:"@odata.type,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`
	UserId                      *string                                      `json:"userId,omitempty"`
}
