package contact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactsProperties struct {
	AntennaConfiguration    *ContactsPropertiesAntennaConfiguration `json:"antennaConfiguration,omitempty"`
	ContactProfile          ResourceReference                       `json:"contactProfile"`
	EndAzimuthDegrees       *float64                                `json:"endAzimuthDegrees,omitempty"`
	EndElevationDegrees     *float64                                `json:"endElevationDegrees,omitempty"`
	ErrorMessage            *string                                 `json:"errorMessage,omitempty"`
	GroundStationName       string                                  `json:"groundStationName"`
	MaximumElevationDegrees *float64                                `json:"maximumElevationDegrees,omitempty"`
	ProvisioningState       *ProvisioningState                      `json:"provisioningState,omitempty"`
	ReservationEndTime      string                                  `json:"reservationEndTime"`
	ReservationStartTime    string                                  `json:"reservationStartTime"`
	RxEndTime               *string                                 `json:"rxEndTime,omitempty"`
	RxStartTime             *string                                 `json:"rxStartTime,omitempty"`
	StartAzimuthDegrees     *float64                                `json:"startAzimuthDegrees,omitempty"`
	StartElevationDegrees   *float64                                `json:"startElevationDegrees,omitempty"`
	Status                  *ContactsStatus                         `json:"status,omitempty"`
	TxEndTime               *string                                 `json:"txEndTime,omitempty"`
	TxStartTime             *string                                 `json:"txStartTime,omitempty"`
}
