package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomerInformation struct {
	CustomQuestionAnswers *[]BookingQuestionAnswer `json:"customQuestionAnswers,omitempty"`
	CustomerId            *string                  `json:"customerId,omitempty"`
	EmailAddress          *string                  `json:"emailAddress,omitempty"`
	Location              *Location                `json:"location,omitempty"`
	Name                  *string                  `json:"name,omitempty"`
	Notes                 *string                  `json:"notes,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
	Phone                 *string                  `json:"phone,omitempty"`
	TimeZone              *string                  `json:"timeZone,omitempty"`
}
