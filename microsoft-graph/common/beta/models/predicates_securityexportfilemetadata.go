package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportFileMetadataOperationPredicate struct {
	DownloadUrl *string
	FileName    *string
	ODataType   *string
	Size        *int64
}

func (p SecurityExportFileMetadataOperationPredicate) Matches(input SecurityExportFileMetadata) bool {

	if p.DownloadUrl != nil && (input.DownloadUrl == nil || *p.DownloadUrl != *input.DownloadUrl) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	return true
}
