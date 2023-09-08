package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationCount struct {
	ODataType          *string                                  `json:"@odata.type,omitempty"`
	RegistrationCount  *int64                                   `json:"registrationCount,omitempty"`
	RegistrationStatus *UserRegistrationCountRegistrationStatus `json:"registrationStatus,omitempty"`
}
