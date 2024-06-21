package testjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestJob struct {
	CreationTime           *string            `json:"creationTime,omitempty"`
	EndTime                *string            `json:"endTime,omitempty"`
	Exception              *string            `json:"exception,omitempty"`
	LastModifiedTime       *string            `json:"lastModifiedTime,omitempty"`
	LastStatusModifiedTime *string            `json:"lastStatusModifiedTime,omitempty"`
	LogActivityTrace       *int64             `json:"logActivityTrace,omitempty"`
	Parameters             *map[string]string `json:"parameters,omitempty"`
	RunOn                  *string            `json:"runOn,omitempty"`
	StartTime              *string            `json:"startTime,omitempty"`
	Status                 *string            `json:"status,omitempty"`
	StatusDetails          *string            `json:"statusDetails,omitempty"`
}
