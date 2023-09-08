package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchSessionBaseOperationPredicate struct {
	CompletionDateTime           *string
	CreationDateTime             *string
	DataStoreId                  *string
	Id                           *string
	LastUpdatedDateTime          *string
	ODataType                    *string
	ProcessingCompletionDateTime *string
	RemainingBlockCount          *int64
	RemainingJobCount            *int64
	StartDateTime                *string
	State                        *string
	TotalBlockCount              *int64
	TotalJobCount                *int64
	UploadCompletionDateTime     *string
}

func (p ExactMatchSessionBaseOperationPredicate) Matches(input ExactMatchSessionBase) bool {

	if p.CompletionDateTime != nil && (input.CompletionDateTime == nil || *p.CompletionDateTime != *input.CompletionDateTime) {
		return false
	}

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.DataStoreId != nil && (input.DataStoreId == nil || *p.DataStoreId != *input.DataStoreId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProcessingCompletionDateTime != nil && (input.ProcessingCompletionDateTime == nil || *p.ProcessingCompletionDateTime != *input.ProcessingCompletionDateTime) {
		return false
	}

	if p.RemainingBlockCount != nil && (input.RemainingBlockCount == nil || *p.RemainingBlockCount != *input.RemainingBlockCount) {
		return false
	}

	if p.RemainingJobCount != nil && (input.RemainingJobCount == nil || *p.RemainingJobCount != *input.RemainingJobCount) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	if p.TotalBlockCount != nil && (input.TotalBlockCount == nil || *p.TotalBlockCount != *input.TotalBlockCount) {
		return false
	}

	if p.TotalJobCount != nil && (input.TotalJobCount == nil || *p.TotalJobCount != *input.TotalJobCount) {
		return false
	}

	if p.UploadCompletionDateTime != nil && (input.UploadCompletionDateTime == nil || *p.UploadCompletionDateTime != *input.UploadCompletionDateTime) {
		return false
	}

	return true
}
