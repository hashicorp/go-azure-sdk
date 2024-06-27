package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationNumber struct {
	Id       *string   `json:"id,omitempty"`
	Required *bool     `json:"required,omitempty"`
	Type     *[]string `json:"type,omitempty"`
}
