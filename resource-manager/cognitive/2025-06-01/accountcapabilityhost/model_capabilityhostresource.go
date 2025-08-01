package accountcapabilityhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityHostResource struct {
	Id         *string        `json:"id,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Properties CapabilityHost `json:"properties"`
	Type       *string        `json:"type,omitempty"`
}
