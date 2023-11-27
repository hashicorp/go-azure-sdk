package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Contact struct {
	AddressMailing *Address `json:"addressMailing,omitempty"`
	Email          string   `json:"email"`
	Fax            *string  `json:"fax,omitempty"`
	JobTitle       *string  `json:"jobTitle,omitempty"`
	NameFirst      string   `json:"nameFirst"`
	NameLast       string   `json:"nameLast"`
	NameMiddle     *string  `json:"nameMiddle,omitempty"`
	Organization   *string  `json:"organization,omitempty"`
	Phone          string   `json:"phone"`
}
