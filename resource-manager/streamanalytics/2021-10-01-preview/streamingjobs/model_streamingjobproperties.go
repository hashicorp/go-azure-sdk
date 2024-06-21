package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingJobProperties struct {
	Cluster                            *ClusterInfo            `json:"cluster,omitempty"`
	CompatibilityLevel                 *CompatibilityLevel     `json:"compatibilityLevel,omitempty"`
	ContentStoragePolicy               *ContentStoragePolicy   `json:"contentStoragePolicy,omitempty"`
	CreatedDate                        *string                 `json:"createdDate,omitempty"`
	DataLocale                         *string                 `json:"dataLocale,omitempty"`
	Etag                               *string                 `json:"etag,omitempty"`
	EventsLateArrivalMaxDelayInSeconds *int64                  `json:"eventsLateArrivalMaxDelayInSeconds,omitempty"`
	EventsOutOfOrderMaxDelayInSeconds  *int64                  `json:"eventsOutOfOrderMaxDelayInSeconds,omitempty"`
	EventsOutOfOrderPolicy             *EventsOutOfOrderPolicy `json:"eventsOutOfOrderPolicy,omitempty"`
	Externals                          *External               `json:"externals,omitempty"`
	Functions                          *[]Function             `json:"functions,omitempty"`
	Inputs                             *[]Input                `json:"inputs,omitempty"`
	JobId                              *string                 `json:"jobId,omitempty"`
	JobState                           *string                 `json:"jobState,omitempty"`
	JobStorageAccount                  *JobStorageAccount      `json:"jobStorageAccount,omitempty"`
	JobType                            *JobType                `json:"jobType,omitempty"`
	LastOutputEventTime                *string                 `json:"lastOutputEventTime,omitempty"`
	OutputErrorPolicy                  *OutputErrorPolicy      `json:"outputErrorPolicy,omitempty"`
	OutputStartMode                    *OutputStartMode        `json:"outputStartMode,omitempty"`
	OutputStartTime                    *string                 `json:"outputStartTime,omitempty"`
	Outputs                            *[]Output               `json:"outputs,omitempty"`
	ProvisioningState                  *string                 `json:"provisioningState,omitempty"`
	Sku                                *Sku                    `json:"sku,omitempty"`
	Transformation                     *Transformation         `json:"transformation,omitempty"`
}
