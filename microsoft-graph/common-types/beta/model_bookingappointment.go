package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointment struct {
	AdditionalInformation    *string                           `json:"additionalInformation,omitempty"`
	AnonymousJoinWebUrl      *string                           `json:"anonymousJoinWebUrl,omitempty"`
	CustomerEmailAddress     *string                           `json:"customerEmailAddress,omitempty"`
	CustomerId               *string                           `json:"customerId,omitempty"`
	CustomerLocation         *Location                         `json:"customerLocation,omitempty"`
	CustomerName             *string                           `json:"customerName,omitempty"`
	CustomerNotes            *string                           `json:"customerNotes,omitempty"`
	CustomerPhone            *string                           `json:"customerPhone,omitempty"`
	CustomerTimeZone         *string                           `json:"customerTimeZone,omitempty"`
	Customers                *[]BookingCustomerInformationBase `json:"customers,omitempty"`
	Duration                 *string                           `json:"duration,omitempty"`
	End                      *DateTimeTimeZone                 `json:"end,omitempty"`
	FilledAttendeesCount     *int64                            `json:"filledAttendeesCount,omitempty"`
	Id                       *string                           `json:"id,omitempty"`
	InvoiceAmount            *float64                          `json:"invoiceAmount,omitempty"`
	InvoiceDate              *DateTimeTimeZone                 `json:"invoiceDate,omitempty"`
	InvoiceId                *string                           `json:"invoiceId,omitempty"`
	InvoiceStatus            *BookingAppointmentInvoiceStatus  `json:"invoiceStatus,omitempty"`
	InvoiceUrl               *string                           `json:"invoiceUrl,omitempty"`
	IsLocationOnline         *bool                             `json:"isLocationOnline,omitempty"`
	JoinWebUrl               *string                           `json:"joinWebUrl,omitempty"`
	MaximumAttendeesCount    *int64                            `json:"maximumAttendeesCount,omitempty"`
	ODataType                *string                           `json:"@odata.type,omitempty"`
	OnlineMeetingUrl         *string                           `json:"onlineMeetingUrl,omitempty"`
	OptOutOfCustomerEmail    *bool                             `json:"optOutOfCustomerEmail,omitempty"`
	PostBuffer               *string                           `json:"postBuffer,omitempty"`
	PreBuffer                *string                           `json:"preBuffer,omitempty"`
	Price                    *float64                          `json:"price,omitempty"`
	PriceType                *BookingAppointmentPriceType      `json:"priceType,omitempty"`
	Reminders                *[]BookingReminder                `json:"reminders,omitempty"`
	SelfServiceAppointmentId *string                           `json:"selfServiceAppointmentId,omitempty"`
	ServiceId                *string                           `json:"serviceId,omitempty"`
	ServiceLocation          *Location                         `json:"serviceLocation,omitempty"`
	ServiceName              *string                           `json:"serviceName,omitempty"`
	ServiceNotes             *string                           `json:"serviceNotes,omitempty"`
	SmsNotificationsEnabled  *bool                             `json:"smsNotificationsEnabled,omitempty"`
	StaffMemberIds           *[]string                         `json:"staffMemberIds,omitempty"`
	Start                    *DateTimeTimeZone                 `json:"start,omitempty"`
}
