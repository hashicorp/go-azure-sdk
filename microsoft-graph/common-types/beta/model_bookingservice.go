package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingService struct {
	AdditionalInformation   *string                         `json:"additionalInformation,omitempty"`
	CustomQuestions         *[]BookingQuestionAssignment    `json:"customQuestions,omitempty"`
	DefaultDuration         *string                         `json:"defaultDuration,omitempty"`
	DefaultLocation         *Location                       `json:"defaultLocation,omitempty"`
	DefaultPrice            *float64                        `json:"defaultPrice,omitempty"`
	DefaultPriceType        *BookingServiceDefaultPriceType `json:"defaultPriceType,omitempty"`
	DefaultReminders        *[]BookingReminder              `json:"defaultReminders,omitempty"`
	Description             *string                         `json:"description,omitempty"`
	DisplayName             *string                         `json:"displayName,omitempty"`
	Id                      *string                         `json:"id,omitempty"`
	IsAnonymousJoinEnabled  *bool                           `json:"isAnonymousJoinEnabled,omitempty"`
	IsHiddenFromCustomers   *bool                           `json:"isHiddenFromCustomers,omitempty"`
	IsLocationOnline        *bool                           `json:"isLocationOnline,omitempty"`
	LanguageTag             *string                         `json:"languageTag,omitempty"`
	MaximumAttendeesCount   *int64                          `json:"maximumAttendeesCount,omitempty"`
	Notes                   *string                         `json:"notes,omitempty"`
	ODataType               *string                         `json:"@odata.type,omitempty"`
	PostBuffer              *string                         `json:"postBuffer,omitempty"`
	PreBuffer               *string                         `json:"preBuffer,omitempty"`
	SchedulingPolicy        *BookingSchedulingPolicy        `json:"schedulingPolicy,omitempty"`
	SmsNotificationsEnabled *bool                           `json:"smsNotificationsEnabled,omitempty"`
	StaffMemberIds          *[]string                       `json:"staffMemberIds,omitempty"`
	WebUrl                  *string                         `json:"webUrl,omitempty"`
}
