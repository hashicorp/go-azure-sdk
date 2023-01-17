package servicerunners

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceRunner struct {
	Id       *string             `json:"id,omitempty"`
	Identity *IdentityProperties `json:"identity,omitempty"`
	Location *string             `json:"location,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Tags     *map[string]string  `json:"tags,omitempty"`
	Type     *string             `json:"type,omitempty"`
}
