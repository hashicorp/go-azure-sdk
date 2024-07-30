package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingBusiness struct {
	Address            *PhysicalAddress          `json:"address,omitempty"`
	Appointments       *[]BookingAppointment     `json:"appointments,omitempty"`
	BusinessHours      *[]BookingWorkHours       `json:"businessHours,omitempty"`
	BusinessType       *string                   `json:"businessType,omitempty"`
	CalendarView       *[]BookingAppointment     `json:"calendarView,omitempty"`
	CustomQuestions    *[]BookingCustomQuestion  `json:"customQuestions,omitempty"`
	Customers          *[]BookingCustomerBase    `json:"customers,omitempty"`
	DefaultCurrencyIso *string                   `json:"defaultCurrencyIso,omitempty"`
	DisplayName        *string                   `json:"displayName,omitempty"`
	Email              *string                   `json:"email,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	IsPublished        *bool                     `json:"isPublished,omitempty"`
	LanguageTag        *string                   `json:"languageTag,omitempty"`
	ODataType          *string                   `json:"@odata.type,omitempty"`
	Phone              *string                   `json:"phone,omitempty"`
	PublicUrl          *string                   `json:"publicUrl,omitempty"`
	SchedulingPolicy   *BookingSchedulingPolicy  `json:"schedulingPolicy,omitempty"`
	Services           *[]BookingService         `json:"services,omitempty"`
	StaffMembers       *[]BookingStaffMemberBase `json:"staffMembers,omitempty"`
	WebSiteUrl         *string                   `json:"webSiteUrl,omitempty"`
}
