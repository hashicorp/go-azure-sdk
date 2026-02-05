package jobschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobScheduleAddParameter struct {
	DisplayName      *string          `json:"displayName,omitempty"`
	Id               string           `json:"id"`
	JobSpecification JobSpecification `json:"jobSpecification"`
	Metadata         *[]MetadataItem  `json:"metadata,omitempty"`
	Schedule         Schedule         `json:"schedule"`
}
