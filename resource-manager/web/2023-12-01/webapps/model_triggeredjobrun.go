package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredJobRun struct {
	Duration   *string                `json:"duration,omitempty"`
	EndTime    *string                `json:"end_time,omitempty"`
	ErrorUrl   *string                `json:"error_url,omitempty"`
	JobName    *string                `json:"job_name,omitempty"`
	OutputUrl  *string                `json:"output_url,omitempty"`
	StartTime  *string                `json:"start_time,omitempty"`
	Status     *TriggeredWebJobStatus `json:"status,omitempty"`
	Trigger    *string                `json:"trigger,omitempty"`
	Url        *string                `json:"url,omitempty"`
	WebJobId   *string                `json:"web_job_id,omitempty"`
	WebJobName *string                `json:"web_job_name,omitempty"`
}
