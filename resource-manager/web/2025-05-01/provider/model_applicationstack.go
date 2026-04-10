package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationStack struct {
	Dependency    *string              `json:"dependency,omitempty"`
	Display       *string              `json:"display,omitempty"`
	Frameworks    *[]ApplicationStack  `json:"frameworks,omitempty"`
	IsDeprecated  *[]ApplicationStack  `json:"isDeprecated,omitempty"`
	MajorVersions *[]StackMajorVersion `json:"majorVersions,omitempty"`
	Name          *string              `json:"name,omitempty"`
}
