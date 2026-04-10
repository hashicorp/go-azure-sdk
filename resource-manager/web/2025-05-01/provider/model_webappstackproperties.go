package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppStackProperties struct {
	DisplayText   *string               `json:"displayText,omitempty"`
	MajorVersions *[]WebAppMajorVersion `json:"majorVersions,omitempty"`
	PreferredOs   *StackPreferredOs     `json:"preferredOs,omitempty"`
	Value         *string               `json:"value,omitempty"`
}
