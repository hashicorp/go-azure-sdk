package customers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerProperties struct {
	DisplayName       *string      `json:"displayName,omitempty"`
	EnabledAzurePlans *[]AzurePlan `json:"enabledAzurePlans,omitempty"`
	Resellers         *[]Reseller  `json:"resellers,omitempty"`
}
