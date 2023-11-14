package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Address struct {
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2,omitempty"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postalCode"`
	State      string  `json:"state"`
}
