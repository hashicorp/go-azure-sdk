package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingAppointmentOperationPredicate struct {
	AdditionalInformation    *string
	AnonymousJoinWebUrl      *string
	CustomerEmailAddress     *string
	CustomerId               *string
	CustomerName             *string
	CustomerNotes            *string
	CustomerPhone            *string
	CustomerTimeZone         *string
	Duration                 *string
	FilledAttendeesCount     *int64
	Id                       *string
	InvoiceId                *string
	InvoiceUrl               *string
	IsLocationOnline         *bool
	JoinWebUrl               *string
	MaximumAttendeesCount    *int64
	ODataType                *string
	OnlineMeetingUrl         *string
	OptOutOfCustomerEmail    *bool
	PostBuffer               *string
	PreBuffer                *string
	SelfServiceAppointmentId *string
	ServiceId                *string
	ServiceName              *string
	ServiceNotes             *string
	SmsNotificationsEnabled  *bool
}

func (p BookingAppointmentOperationPredicate) Matches(input BookingAppointment) bool {

	if p.AdditionalInformation != nil && (input.AdditionalInformation == nil || *p.AdditionalInformation != *input.AdditionalInformation) {
		return false
	}

	if p.AnonymousJoinWebUrl != nil && (input.AnonymousJoinWebUrl == nil || *p.AnonymousJoinWebUrl != *input.AnonymousJoinWebUrl) {
		return false
	}

	if p.CustomerEmailAddress != nil && (input.CustomerEmailAddress == nil || *p.CustomerEmailAddress != *input.CustomerEmailAddress) {
		return false
	}

	if p.CustomerId != nil && (input.CustomerId == nil || *p.CustomerId != *input.CustomerId) {
		return false
	}

	if p.CustomerName != nil && (input.CustomerName == nil || *p.CustomerName != *input.CustomerName) {
		return false
	}

	if p.CustomerNotes != nil && (input.CustomerNotes == nil || *p.CustomerNotes != *input.CustomerNotes) {
		return false
	}

	if p.CustomerPhone != nil && (input.CustomerPhone == nil || *p.CustomerPhone != *input.CustomerPhone) {
		return false
	}

	if p.CustomerTimeZone != nil && (input.CustomerTimeZone == nil || *p.CustomerTimeZone != *input.CustomerTimeZone) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.FilledAttendeesCount != nil && (input.FilledAttendeesCount == nil || *p.FilledAttendeesCount != *input.FilledAttendeesCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InvoiceId != nil && (input.InvoiceId == nil || *p.InvoiceId != *input.InvoiceId) {
		return false
	}

	if p.InvoiceUrl != nil && (input.InvoiceUrl == nil || *p.InvoiceUrl != *input.InvoiceUrl) {
		return false
	}

	if p.IsLocationOnline != nil && (input.IsLocationOnline == nil || *p.IsLocationOnline != *input.IsLocationOnline) {
		return false
	}

	if p.JoinWebUrl != nil && (input.JoinWebUrl == nil || *p.JoinWebUrl != *input.JoinWebUrl) {
		return false
	}

	if p.MaximumAttendeesCount != nil && (input.MaximumAttendeesCount == nil || *p.MaximumAttendeesCount != *input.MaximumAttendeesCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnlineMeetingUrl != nil && (input.OnlineMeetingUrl == nil || *p.OnlineMeetingUrl != *input.OnlineMeetingUrl) {
		return false
	}

	if p.OptOutOfCustomerEmail != nil && (input.OptOutOfCustomerEmail == nil || *p.OptOutOfCustomerEmail != *input.OptOutOfCustomerEmail) {
		return false
	}

	if p.PostBuffer != nil && (input.PostBuffer == nil || *p.PostBuffer != *input.PostBuffer) {
		return false
	}

	if p.PreBuffer != nil && (input.PreBuffer == nil || *p.PreBuffer != *input.PreBuffer) {
		return false
	}

	if p.SelfServiceAppointmentId != nil && (input.SelfServiceAppointmentId == nil || *p.SelfServiceAppointmentId != *input.SelfServiceAppointmentId) {
		return false
	}

	if p.ServiceId != nil && (input.ServiceId == nil || *p.ServiceId != *input.ServiceId) {
		return false
	}

	if p.ServiceName != nil && (input.ServiceName == nil || *p.ServiceName != *input.ServiceName) {
		return false
	}

	if p.ServiceNotes != nil && (input.ServiceNotes == nil || *p.ServiceNotes != *input.ServiceNotes) {
		return false
	}

	if p.SmsNotificationsEnabled != nil && (input.SmsNotificationsEnabled == nil || *p.SmsNotificationsEnabled != *input.SmsNotificationsEnabled) {
		return false
	}

	return true
}
