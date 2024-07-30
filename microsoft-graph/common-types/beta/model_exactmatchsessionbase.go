package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchSessionBase struct {
	CompletionDateTime           *string              `json:"completionDateTime,omitempty"`
	CreationDateTime             *string              `json:"creationDateTime,omitempty"`
	DataStoreId                  *string              `json:"dataStoreId,omitempty"`
	Error                        *ClassificationError `json:"error,omitempty"`
	Id                           *string              `json:"id,omitempty"`
	LastUpdatedDateTime          *string              `json:"lastUpdatedDateTime,omitempty"`
	ODataType                    *string              `json:"@odata.type,omitempty"`
	ProcessingCompletionDateTime *string              `json:"processingCompletionDateTime,omitempty"`
	RemainingBlockCount          *int64               `json:"remainingBlockCount,omitempty"`
	RemainingJobCount            *int64               `json:"remainingJobCount,omitempty"`
	StartDateTime                *string              `json:"startDateTime,omitempty"`
	State                        *string              `json:"state,omitempty"`
	TotalBlockCount              *int64               `json:"totalBlockCount,omitempty"`
	TotalJobCount                *int64               `json:"totalJobCount,omitempty"`
	UploadCompletionDateTime     *string              `json:"uploadCompletionDateTime,omitempty"`
}
