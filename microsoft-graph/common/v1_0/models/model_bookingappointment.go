package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointment struct {
	AdditionalInformation    *string                           `json:"additionalInformation,omitempty"`
	AnonymousJoinWebUrl      *string                           `json:"anonymousJoinWebUrl,omitempty"`
	CustomerTimeZone         *string                           `json:"customerTimeZone,omitempty"`
	Customers                *[]BookingCustomerInformationBase `json:"customers,omitempty"`
	Duration                 *string                           `json:"duration,omitempty"`
	EndDateTime              *DateTimeTimeZone                 `json:"endDateTime,omitempty"`
	FilledAttendeesCount     *int64                            `json:"filledAttendeesCount,omitempty"`
	Id                       *string                           `json:"id,omitempty"`
	IsLocationOnline         *bool                             `json:"isLocationOnline,omitempty"`
	JoinWebUrl               *string                           `json:"joinWebUrl,omitempty"`
	MaximumAttendeesCount    *int64                            `json:"maximumAttendeesCount,omitempty"`
	ODataType                *string                           `json:"@odata.type,omitempty"`
	OptOutOfCustomerEmail    *bool                             `json:"optOutOfCustomerEmail,omitempty"`
	PostBuffer               *string                           `json:"postBuffer,omitempty"`
	PreBuffer                *string                           `json:"preBuffer,omitempty"`
	PriceType                *BookingAppointmentPriceType      `json:"priceType,omitempty"`
	Reminders                *[]BookingReminder                `json:"reminders,omitempty"`
	SelfServiceAppointmentId *string                           `json:"selfServiceAppointmentId,omitempty"`
	ServiceId                *string                           `json:"serviceId,omitempty"`
	ServiceLocation          *Location                         `json:"serviceLocation,omitempty"`
	ServiceName              *string                           `json:"serviceName,omitempty"`
	ServiceNotes             *string                           `json:"serviceNotes,omitempty"`
	SmsNotificationsEnabled  *bool                             `json:"smsNotificationsEnabled,omitempty"`
	StaffMemberIds           *[]string                         `json:"staffMemberIds,omitempty"`
	StartDateTime            *DateTimeTimeZone                 `json:"startDateTime,omitempty"`
}
