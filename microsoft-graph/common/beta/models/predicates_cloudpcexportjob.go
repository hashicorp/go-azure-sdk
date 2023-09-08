package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobOperationPredicate struct {
	ExpirationDateTime *string
	ExportUrl          *string
	Filter             *string
	Format             *string
	Id                 *string
	ODataType          *string
	RequestDateTime    *string
}

func (p CloudPCExportJobOperationPredicate) Matches(input CloudPCExportJob) bool {

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.ExportUrl != nil && (input.ExportUrl == nil || *p.ExportUrl != *input.ExportUrl) {
		return false
	}

	if p.Filter != nil && (input.Filter == nil || *p.Filter != *input.Filter) {
		return false
	}

	if p.Format != nil && (input.Format == nil || *p.Format != *input.Format) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestDateTime != nil && (input.RequestDateTime == nil || *p.RequestDateTime != *input.RequestDateTime) {
		return false
	}

	return true
}
