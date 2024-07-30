package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTroubleshootingEvent struct {
	AdditionalInformation       *[]KeyValuePair                                `json:"additionalInformation,omitempty"`
	CorrelationId               *string                                        `json:"correlationId,omitempty"`
	DeviceId                    *string                                        `json:"deviceId,omitempty"`
	EnrollmentType              *EnrollmentTroubleshootingEventEnrollmentType  `json:"enrollmentType,omitempty"`
	EventDateTime               *string                                        `json:"eventDateTime,omitempty"`
	EventName                   *string                                        `json:"eventName,omitempty"`
	FailureCategory             *EnrollmentTroubleshootingEventFailureCategory `json:"failureCategory,omitempty"`
	FailureReason               *string                                        `json:"failureReason,omitempty"`
	Id                          *string                                        `json:"id,omitempty"`
	ManagedDeviceIdentifier     *string                                        `json:"managedDeviceIdentifier,omitempty"`
	ODataType                   *string                                        `json:"@odata.type,omitempty"`
	OperatingSystem             *string                                        `json:"operatingSystem,omitempty"`
	OsVersion                   *string                                        `json:"osVersion,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails   `json:"troubleshootingErrorDetails,omitempty"`
	UserId                      *string                                        `json:"userId,omitempty"`
}
