package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchSession struct {
	Checksum                     *string                `json:"checksum,omitempty"`
	CompletionDateTime           *string                `json:"completionDateTime,omitempty"`
	CreationDateTime             *string                `json:"creationDateTime,omitempty"`
	DataStoreId                  *string                `json:"dataStoreId,omitempty"`
	DataUploadURI                *string                `json:"dataUploadURI,omitempty"`
	Error                        *ClassificationError   `json:"error,omitempty"`
	Fields                       *[]string              `json:"fields,omitempty"`
	FileName                     *string                `json:"fileName,omitempty"`
	Id                           *string                `json:"id,omitempty"`
	LastUpdatedDateTime          *string                `json:"lastUpdatedDateTime,omitempty"`
	ODataType                    *string                `json:"@odata.type,omitempty"`
	ProcessingCompletionDateTime *string                `json:"processingCompletionDateTime,omitempty"`
	RemainingBlockCount          *int64                 `json:"remainingBlockCount,omitempty"`
	RemainingJobCount            *int64                 `json:"remainingJobCount,omitempty"`
	RowsPerBlock                 *int64                 `json:"rowsPerBlock,omitempty"`
	Salt                         *string                `json:"salt,omitempty"`
	StartDateTime                *string                `json:"startDateTime,omitempty"`
	State                        *string                `json:"state,omitempty"`
	TotalBlockCount              *int64                 `json:"totalBlockCount,omitempty"`
	TotalJobCount                *int64                 `json:"totalJobCount,omitempty"`
	UploadAgent                  *ExactMatchUploadAgent `json:"uploadAgent,omitempty"`
	UploadAgentId                *string                `json:"uploadAgentId,omitempty"`
	UploadCompletionDateTime     *string                `json:"uploadCompletionDateTime,omitempty"`
}
