package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionResponse struct {
	EnrolledByUser               *string                            `json:"enrolledByUser,omitempty"`
	ErrorCode                    *int64                             `json:"errorCode,omitempty"`
	ExpirationDateTimeUTC        *string                            `json:"expirationDateTimeUTC,omitempty"`
	Id                           *string                            `json:"id,omitempty"`
	InitiatedByUserPrincipalName *string                            `json:"initiatedByUserPrincipalName,omitempty"`
	ManagedDeviceId              *string                            `json:"managedDeviceId,omitempty"`
	ODataType                    *string                            `json:"@odata.type,omitempty"`
	ReceivedDateTimeUTC          *string                            `json:"receivedDateTimeUTC,omitempty"`
	RequestedDateTimeUTC         *string                            `json:"requestedDateTimeUTC,omitempty"`
	Status                       *DeviceLogCollectionResponseStatus `json:"status,omitempty"`
}
