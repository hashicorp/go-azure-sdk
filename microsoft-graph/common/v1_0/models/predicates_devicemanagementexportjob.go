package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobOperationPredicate struct {
	ExpirationDateTime *string
	Filter             *string
	Id                 *string
	ODataType          *string
	ReportName         *string
	RequestDateTime    *string
	SnapshotId         *string
	Url                *string
}

func (p DeviceManagementExportJobOperationPredicate) Matches(input DeviceManagementExportJob) bool {

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Filter != nil && (input.Filter == nil || *p.Filter != *input.Filter) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReportName != nil && (input.ReportName == nil || *p.ReportName != *input.ReportName) {
		return false
	}

	if p.RequestDateTime != nil && (input.RequestDateTime == nil || *p.RequestDateTime != *input.RequestDateTime) {
		return false
	}

	if p.SnapshotId != nil && (input.SnapshotId == nil || *p.SnapshotId != *input.SnapshotId) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
