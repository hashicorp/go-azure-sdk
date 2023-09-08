package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppRestartSettingsOperationPredicate struct {
	CountdownDisplayBeforeRestartInMinutes     *int64
	GracePeriodInMinutes                       *int64
	ODataType                                  *string
	RestartNotificationSnoozeDurationInMinutes *int64
}

func (p WinGetAppRestartSettingsOperationPredicate) Matches(input WinGetAppRestartSettings) bool {

	if p.CountdownDisplayBeforeRestartInMinutes != nil && (input.CountdownDisplayBeforeRestartInMinutes == nil || *p.CountdownDisplayBeforeRestartInMinutes != *input.CountdownDisplayBeforeRestartInMinutes) {
		return false
	}

	if p.GracePeriodInMinutes != nil && (input.GracePeriodInMinutes == nil || *p.GracePeriodInMinutes != *input.GracePeriodInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RestartNotificationSnoozeDurationInMinutes != nil && (input.RestartNotificationSnoozeDurationInMinutes == nil || *p.RestartNotificationSnoozeDurationInMinutes != *input.RestartNotificationSnoozeDurationInMinutes) {
		return false
	}

	return true
}
