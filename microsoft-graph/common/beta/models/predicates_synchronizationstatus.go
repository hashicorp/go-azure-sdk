package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationStatusOperationPredicate struct {
	CountSuccessiveCompleteFailures *int64
	EscrowsPruned                   *bool
	ODataType                       *string
	SteadyStateFirstAchievedTime    *string
	SteadyStateLastAchievedTime     *string
	TroubleshootingUrl              *string
}

func (p SynchronizationStatusOperationPredicate) Matches(input SynchronizationStatus) bool {

	if p.CountSuccessiveCompleteFailures != nil && (input.CountSuccessiveCompleteFailures == nil || *p.CountSuccessiveCompleteFailures != *input.CountSuccessiveCompleteFailures) {
		return false
	}

	if p.EscrowsPruned != nil && (input.EscrowsPruned == nil || *p.EscrowsPruned != *input.EscrowsPruned) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SteadyStateFirstAchievedTime != nil && (input.SteadyStateFirstAchievedTime == nil || *p.SteadyStateFirstAchievedTime != *input.SteadyStateFirstAchievedTime) {
		return false
	}

	if p.SteadyStateLastAchievedTime != nil && (input.SteadyStateLastAchievedTime == nil || *p.SteadyStateLastAchievedTime != *input.SteadyStateLastAchievedTime) {
		return false
	}

	if p.TroubleshootingUrl != nil && (input.TroubleshootingUrl == nil || *p.TroubleshootingUrl != *input.TroubleshootingUrl) {
		return false
	}

	return true
}
