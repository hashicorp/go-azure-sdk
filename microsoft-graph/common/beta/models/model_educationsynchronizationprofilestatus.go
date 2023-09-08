package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfileStatus struct {
	ErrorCount                  *int64                                       `json:"errorCount,omitempty"`
	Id                          *string                                      `json:"id,omitempty"`
	LastActivityDateTime        *string                                      `json:"lastActivityDateTime,omitempty"`
	LastSynchronizationDateTime *string                                      `json:"lastSynchronizationDateTime,omitempty"`
	ODataType                   *string                                      `json:"@odata.type,omitempty"`
	Status                      *EducationSynchronizationProfileStatusStatus `json:"status,omitempty"`
	StatusMessage               *string                                      `json:"statusMessage,omitempty"`
}
