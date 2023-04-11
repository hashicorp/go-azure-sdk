package informationprotectionpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionKeyword struct {
	CanBeNumeric *bool   `json:"canBeNumeric,omitempty"`
	Custom       *bool   `json:"custom,omitempty"`
	Excluded     *bool   `json:"excluded,omitempty"`
	Pattern      *string `json:"pattern,omitempty"`
}
