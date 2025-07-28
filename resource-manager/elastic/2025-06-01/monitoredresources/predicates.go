package monitoredresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoredResourceOperationPredicate struct {
	Id                  *string
	ReasonForLogsStatus *string
}

func (p MonitoredResourceOperationPredicate) Matches(input MonitoredResource) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ReasonForLogsStatus != nil && (input.ReasonForLogsStatus == nil || *p.ReasonForLogsStatus != *input.ReasonForLogsStatus) {
		return false
	}

	return true
}
