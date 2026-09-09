package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppMajorVersion struct {
	DisplayText   *string               `json:"displayText,omitempty"`
	MinorVersions *[]WebAppMinorVersion `json:"minorVersions,omitempty"`
	Value         *string               `json:"value,omitempty"`
}
