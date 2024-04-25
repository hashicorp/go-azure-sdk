package configurationnames

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationName struct {
	Description *string `json:"description,omitempty"`
	Required    *bool   `json:"required,omitempty"`
	Value       *string `json:"value,omitempty"`
}
