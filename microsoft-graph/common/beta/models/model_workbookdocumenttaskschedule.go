package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookDocumentTaskSchedule struct {
	DueDateTime   *string `json:"dueDateTime,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}
