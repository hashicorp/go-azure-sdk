package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchSessionOperationPredicate struct {
	Checksum                     *string
	CompletionDateTime           *string
	CreationDateTime             *string
	DataStoreId                  *string
	DataUploadURI                *string
	FileName                     *string
	Id                           *string
	LastUpdatedDateTime          *string
	ODataType                    *string
	ProcessingCompletionDateTime *string
	RemainingBlockCount          *int64
	RemainingJobCount            *int64
	RowsPerBlock                 *int64
	Salt                         *string
	StartDateTime                *string
	State                        *string
	TotalBlockCount              *int64
	TotalJobCount                *int64
	UploadAgentId                *string
	UploadCompletionDateTime     *string
}

func (p ExactMatchSessionOperationPredicate) Matches(input ExactMatchSession) bool {

	if p.Checksum != nil && (input.Checksum == nil || *p.Checksum != *input.Checksum) {
		return false
	}

	if p.CompletionDateTime != nil && (input.CompletionDateTime == nil || *p.CompletionDateTime != *input.CompletionDateTime) {
		return false
	}

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.DataStoreId != nil && (input.DataStoreId == nil || *p.DataStoreId != *input.DataStoreId) {
		return false
	}

	if p.DataUploadURI != nil && (input.DataUploadURI == nil || *p.DataUploadURI != *input.DataUploadURI) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
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

	if p.RowsPerBlock != nil && (input.RowsPerBlock == nil || *p.RowsPerBlock != *input.RowsPerBlock) {
		return false
	}

	if p.Salt != nil && (input.Salt == nil || *p.Salt != *input.Salt) {
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

	if p.UploadAgentId != nil && (input.UploadAgentId == nil || *p.UploadAgentId != *input.UploadAgentId) {
		return false
	}

	if p.UploadCompletionDateTime != nil && (input.UploadCompletionDateTime == nil || *p.UploadCompletionDateTime != *input.UploadCompletionDateTime) {
		return false
	}

	return true
}
