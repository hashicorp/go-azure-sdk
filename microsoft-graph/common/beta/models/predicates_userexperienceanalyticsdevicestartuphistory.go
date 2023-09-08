package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate struct {
	CoreBootTimeInMs          *int64
	CoreLoginTimeInMs         *int64
	DeviceId                  *string
	FeatureUpdateBootTimeInMs *int64
	GroupPolicyBootTimeInMs   *int64
	GroupPolicyLoginTimeInMs  *int64
	Id                        *string
	IsFeatureUpdate           *bool
	IsFirstLogin              *bool
	ODataType                 *string
	OperatingSystemVersion    *string
	ResponsiveDesktopTimeInMs *int64
	RestartFaultBucket        *string
	RestartStopCode           *string
	StartTime                 *string
	TotalBootTimeInMs         *int64
	TotalLoginTimeInMs        *int64
}

func (p UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate) Matches(input UserExperienceAnalyticsDeviceStartupHistory) bool {

	if p.CoreBootTimeInMs != nil && (input.CoreBootTimeInMs == nil || *p.CoreBootTimeInMs != *input.CoreBootTimeInMs) {
		return false
	}

	if p.CoreLoginTimeInMs != nil && (input.CoreLoginTimeInMs == nil || *p.CoreLoginTimeInMs != *input.CoreLoginTimeInMs) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.FeatureUpdateBootTimeInMs != nil && (input.FeatureUpdateBootTimeInMs == nil || *p.FeatureUpdateBootTimeInMs != *input.FeatureUpdateBootTimeInMs) {
		return false
	}

	if p.GroupPolicyBootTimeInMs != nil && (input.GroupPolicyBootTimeInMs == nil || *p.GroupPolicyBootTimeInMs != *input.GroupPolicyBootTimeInMs) {
		return false
	}

	if p.GroupPolicyLoginTimeInMs != nil && (input.GroupPolicyLoginTimeInMs == nil || *p.GroupPolicyLoginTimeInMs != *input.GroupPolicyLoginTimeInMs) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsFeatureUpdate != nil && (input.IsFeatureUpdate == nil || *p.IsFeatureUpdate != *input.IsFeatureUpdate) {
		return false
	}

	if p.IsFirstLogin != nil && (input.IsFirstLogin == nil || *p.IsFirstLogin != *input.IsFirstLogin) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystemVersion != nil && (input.OperatingSystemVersion == nil || *p.OperatingSystemVersion != *input.OperatingSystemVersion) {
		return false
	}

	if p.ResponsiveDesktopTimeInMs != nil && (input.ResponsiveDesktopTimeInMs == nil || *p.ResponsiveDesktopTimeInMs != *input.ResponsiveDesktopTimeInMs) {
		return false
	}

	if p.RestartFaultBucket != nil && (input.RestartFaultBucket == nil || *p.RestartFaultBucket != *input.RestartFaultBucket) {
		return false
	}

	if p.RestartStopCode != nil && (input.RestartStopCode == nil || *p.RestartStopCode != *input.RestartStopCode) {
		return false
	}

	if p.StartTime != nil && (input.StartTime == nil || *p.StartTime != *input.StartTime) {
		return false
	}

	if p.TotalBootTimeInMs != nil && (input.TotalBootTimeInMs == nil || *p.TotalBootTimeInMs != *input.TotalBootTimeInMs) {
		return false
	}

	if p.TotalLoginTimeInMs != nil && (input.TotalLoginTimeInMs == nil || *p.TotalLoginTimeInMs != *input.TotalLoginTimeInMs) {
		return false
	}

	return true
}
