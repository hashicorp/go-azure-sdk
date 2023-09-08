package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatusOperationPredicate struct {
	CancelRequested            *bool
	CompleteOrCanceledDateTime *string
	LastUpdatedDateTime        *string
	ODataType                  *string
	TotalAwaitingInstall       *int64
	TotalCanceled              *int64
	TotalCreated               *int64
	TotalDevices               *int64
	TotalDownloading           *int64
	TotalFailedDownload        *int64
	TotalFailedInstall         *int64
	TotalScheduled             *int64
	TotalSucceededInstall      *int64
	TotalUnknown               *int64
}

func (p ZebraFotaDeploymentStatusOperationPredicate) Matches(input ZebraFotaDeploymentStatus) bool {

	if p.CancelRequested != nil && (input.CancelRequested == nil || *p.CancelRequested != *input.CancelRequested) {
		return false
	}

	if p.CompleteOrCanceledDateTime != nil && (input.CompleteOrCanceledDateTime == nil || *p.CompleteOrCanceledDateTime != *input.CompleteOrCanceledDateTime) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalAwaitingInstall != nil && (input.TotalAwaitingInstall == nil || *p.TotalAwaitingInstall != *input.TotalAwaitingInstall) {
		return false
	}

	if p.TotalCanceled != nil && (input.TotalCanceled == nil || *p.TotalCanceled != *input.TotalCanceled) {
		return false
	}

	if p.TotalCreated != nil && (input.TotalCreated == nil || *p.TotalCreated != *input.TotalCreated) {
		return false
	}

	if p.TotalDevices != nil && (input.TotalDevices == nil || *p.TotalDevices != *input.TotalDevices) {
		return false
	}

	if p.TotalDownloading != nil && (input.TotalDownloading == nil || *p.TotalDownloading != *input.TotalDownloading) {
		return false
	}

	if p.TotalFailedDownload != nil && (input.TotalFailedDownload == nil || *p.TotalFailedDownload != *input.TotalFailedDownload) {
		return false
	}

	if p.TotalFailedInstall != nil && (input.TotalFailedInstall == nil || *p.TotalFailedInstall != *input.TotalFailedInstall) {
		return false
	}

	if p.TotalScheduled != nil && (input.TotalScheduled == nil || *p.TotalScheduled != *input.TotalScheduled) {
		return false
	}

	if p.TotalSucceededInstall != nil && (input.TotalSucceededInstall == nil || *p.TotalSucceededInstall != *input.TotalSucceededInstall) {
		return false
	}

	if p.TotalUnknown != nil && (input.TotalUnknown == nil || *p.TotalUnknown != *input.TotalUnknown) {
		return false
	}

	return true
}
