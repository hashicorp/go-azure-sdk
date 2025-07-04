package disks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessUri struct {
	StartTime  *string              `json:"startTime,omitempty"`
	EndTime    *string              `json:"endTime,omitempty"`
	Status     *string              `json:"status,omitempty"`
	Name       *string              `json:"name,omitempty"`
	Properties *AccessUriProperties `json:"properties,omitempty"`
}
