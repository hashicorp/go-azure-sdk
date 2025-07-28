package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageResult struct {
	Id         *string          `json:"id,omitempty"`
	Name       *UsageName       `json:"name,omitempty"`
	Properties *UsageProperties `json:"properties,omitempty"`
}
