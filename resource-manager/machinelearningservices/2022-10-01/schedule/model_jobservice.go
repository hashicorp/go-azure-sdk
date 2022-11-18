package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobService struct {
	Endpoint       *string            `json:"endpoint,omitempty"`
	ErrorMessage   *string            `json:"errorMessage,omitempty"`
	JobServiceType *string            `json:"jobServiceType,omitempty"`
	Port           *int64             `json:"port,omitempty"`
	Properties     *map[string]string `json:"properties,omitempty"`
	Status         *string            `json:"status,omitempty"`
}
