package jobschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobSchedulePatchParameter struct {
	JobSpecification *JobSpecification `json:"jobSpecification,omitempty"`
	Metadata         *[]MetadataItem   `json:"metadata,omitempty"`
	Schedule         *Schedule         `json:"schedule,omitempty"`
}
