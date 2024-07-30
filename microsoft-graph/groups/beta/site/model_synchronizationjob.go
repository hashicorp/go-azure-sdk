package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJob struct {
	BulkUpload                 *BulkUpload              `json:"bulkUpload,omitempty"`
	Id                         *string                  `json:"id,omitempty"`
	ODataType                  *string                  `json:"@odata.type,omitempty"`
	Schedule                   *SynchronizationSchedule `json:"schedule,omitempty"`
	Schema                     *SynchronizationSchema   `json:"schema,omitempty"`
	Status                     *SynchronizationStatus   `json:"status,omitempty"`
	SynchronizationJobSettings *[]KeyValuePair          `json:"synchronizationJobSettings,omitempty"`
	TemplateId                 *string                  `json:"templateId,omitempty"`
}
