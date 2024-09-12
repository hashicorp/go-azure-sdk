package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BookingService{}

type BookingService struct {
	// Additional information that is sent to the customer when an appointment is confirmed.
	AdditionalInformation nullable.Type[string] `json:"additionalInformation,omitempty"`

	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains the set of custom questions associated with a particular service.
	CustomQuestions *[]BookingQuestionAssignment `json:"customQuestions,omitempty"`

	// The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example,
	// P11D23H59M59.999999999999S.
	DefaultDuration *string `json:"defaultDuration,omitempty"`

	// The default physical location for the service.
	DefaultLocation Location `json:"defaultLocation"`

	// Represents the type of pricing of a booking service.
	DefaultPriceType *BookingPriceType `json:"defaultPriceType,omitempty"`

	// The default set of reminders for an appointment of this service. The value of this property is available only when
	// reading this bookingService by its ID.
	DefaultReminders *[]BookingReminder `json:"defaultReminders,omitempty"`

	// A text description for the service.
	Description nullable.Type[string] `json:"description,omitempty"`

	// A service name.
	DisplayName *string `json:"displayName,omitempty"`

	// True if the URL to join the appointment anonymously (anonymousJoinWebUrl) will be generated for the appointment
	// booked for this service.
	IsAnonymousJoinEnabled *bool `json:"isAnonymousJoinEnabled,omitempty"`

	IsCustomerAllowedToManageBooking nullable.Type[bool] `json:"isCustomerAllowedToManageBooking,omitempty"`

	// True means this service is not available to customers for booking.
	IsHiddenFromCustomers *bool `json:"isHiddenFromCustomers,omitempty"`

	// True indicates that the appointments for the service will be held online. Default value is false.
	IsLocationOnline *bool `json:"isLocationOnline,omitempty"`

	// The language of the self-service booking page.
	LanguageTag *string `json:"languageTag,omitempty"`

	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass
	// valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer
	// operation.
	MaximumAttendeesCount *int64 `json:"maximumAttendeesCount,omitempty"`

	// Additional information about this service.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The time to buffer after an appointment for this service ends, and before the next customer appointment can be
	// booked.
	PostBuffer *string `json:"postBuffer,omitempty"`

	// The time to buffer before an appointment for this service can start.
	PreBuffer *string `json:"preBuffer,omitempty"`

	// The set of policies that determine how appointments for this type of service should be created and managed.
	SchedulingPolicy *BookingSchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is
	// false.
	SmsNotificationsEnabled *bool `json:"smsNotificationsEnabled,omitempty"`

	// Represents those staff members who provide this service.
	StaffMemberIds *[]string `json:"staffMemberIds,omitempty"`

	// The URL a customer uses to access the service.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BookingService) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingService{}

func (s BookingService) MarshalJSON() ([]byte, error) {
	type wrapper BookingService
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingService: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingService: %+v", err)
	}

	delete(decoded, "webUrl")
	decoded["@odata.type"] = "#microsoft.graph.bookingService"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingService: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BookingService{}

func (s *BookingService) UnmarshalJSON(bytes []byte) error {
	type alias BookingService
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BookingService: %+v", err)
	}

	s.AdditionalInformation = decoded.AdditionalInformation
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomQuestions = decoded.CustomQuestions
	s.DefaultDuration = decoded.DefaultDuration
	s.DefaultPriceType = decoded.DefaultPriceType
	s.DefaultReminders = decoded.DefaultReminders
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IsAnonymousJoinEnabled = decoded.IsAnonymousJoinEnabled
	s.IsCustomerAllowedToManageBooking = decoded.IsCustomerAllowedToManageBooking
	s.IsHiddenFromCustomers = decoded.IsHiddenFromCustomers
	s.IsLocationOnline = decoded.IsLocationOnline
	s.LanguageTag = decoded.LanguageTag
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.MaximumAttendeesCount = decoded.MaximumAttendeesCount
	s.Notes = decoded.Notes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PostBuffer = decoded.PostBuffer
	s.PreBuffer = decoded.PreBuffer
	s.SchedulingPolicy = decoded.SchedulingPolicy
	s.SmsNotificationsEnabled = decoded.SmsNotificationsEnabled
	s.StaffMemberIds = decoded.StaffMemberIds
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BookingService into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultLocation"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefaultLocation' for 'BookingService': %+v", err)
		}
		s.DefaultLocation = impl
	}
	return nil
}
