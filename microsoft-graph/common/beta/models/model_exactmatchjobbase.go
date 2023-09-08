package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchJobBase struct {
	CompletionDateTime  *string              `json:"completionDateTime,omitempty"`
	CreationDateTime    *string              `json:"creationDateTime,omitempty"`
	Error               *ClassificationError `json:"error,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	LastUpdatedDateTime *string              `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	StartDateTime       *string              `json:"startDateTime,omitempty"`
}
