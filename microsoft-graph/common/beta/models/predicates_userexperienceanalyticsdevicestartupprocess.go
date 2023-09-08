package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupProcessOperationPredicate struct {
	Id                *string
	ManagedDeviceId   *string
	ODataType         *string
	ProcessName       *string
	ProductName       *string
	Publisher         *string
	StartupImpactInMs *int64
}

func (p UserExperienceAnalyticsDeviceStartupProcessOperationPredicate) Matches(input UserExperienceAnalyticsDeviceStartupProcess) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProcessName != nil && (input.ProcessName == nil || *p.ProcessName != *input.ProcessName) {
		return false
	}

	if p.ProductName != nil && (input.ProductName == nil || *p.ProductName != *input.ProductName) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.StartupImpactInMs != nil && (input.StartupImpactInMs == nil || *p.StartupImpactInMs != *input.StartupImpactInMs) {
		return false
	}

	return true
}
