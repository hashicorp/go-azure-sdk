package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationOperationPredicate struct {
	AllowWindows11Upgrade                   *bool
	CreatedDateTime                         *string
	DeadlineForFeatureUpdatesInDays         *int64
	DeadlineForQualityUpdatesInDays         *int64
	DeadlineGracePeriodInDays               *int64
	Description                             *string
	DisplayName                             *string
	DriversExcluded                         *bool
	EngagedRestartDeadlineInDays            *int64
	EngagedRestartSnoozeScheduleInDays      *int64
	EngagedRestartTransitionScheduleInDays  *int64
	FeatureUpdatesDeferralPeriodInDays      *int64
	FeatureUpdatesPauseExpiryDateTime       *string
	FeatureUpdatesPauseStartDate            *string
	FeatureUpdatesPaused                    *bool
	FeatureUpdatesRollbackStartDateTime     *string
	FeatureUpdatesRollbackWindowInDays      *int64
	FeatureUpdatesWillBeRolledBack          *bool
	Id                                      *string
	LastModifiedDateTime                    *string
	MicrosoftUpdateServiceAllowed           *bool
	ODataType                               *string
	PostponeRebootUntilAfterDeadline        *bool
	QualityUpdatesDeferralPeriodInDays      *int64
	QualityUpdatesPauseExpiryDateTime       *string
	QualityUpdatesPauseStartDate            *string
	QualityUpdatesPaused                    *bool
	QualityUpdatesRollbackStartDateTime     *string
	QualityUpdatesWillBeRolledBack          *bool
	ScheduleImminentRestartWarningInMinutes *int64
	ScheduleRestartWarningInHours           *int64
	SkipChecksBeforeRestart                 *bool
	Version                                 *int64
}

func (p WindowsUpdateForBusinessConfigurationOperationPredicate) Matches(input WindowsUpdateForBusinessConfiguration) bool {

	if p.AllowWindows11Upgrade != nil && (input.AllowWindows11Upgrade == nil || *p.AllowWindows11Upgrade != *input.AllowWindows11Upgrade) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeadlineForFeatureUpdatesInDays != nil && (input.DeadlineForFeatureUpdatesInDays == nil || *p.DeadlineForFeatureUpdatesInDays != *input.DeadlineForFeatureUpdatesInDays) {
		return false
	}

	if p.DeadlineForQualityUpdatesInDays != nil && (input.DeadlineForQualityUpdatesInDays == nil || *p.DeadlineForQualityUpdatesInDays != *input.DeadlineForQualityUpdatesInDays) {
		return false
	}

	if p.DeadlineGracePeriodInDays != nil && (input.DeadlineGracePeriodInDays == nil || *p.DeadlineGracePeriodInDays != *input.DeadlineGracePeriodInDays) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DriversExcluded != nil && (input.DriversExcluded == nil || *p.DriversExcluded != *input.DriversExcluded) {
		return false
	}

	if p.EngagedRestartDeadlineInDays != nil && (input.EngagedRestartDeadlineInDays == nil || *p.EngagedRestartDeadlineInDays != *input.EngagedRestartDeadlineInDays) {
		return false
	}

	if p.EngagedRestartSnoozeScheduleInDays != nil && (input.EngagedRestartSnoozeScheduleInDays == nil || *p.EngagedRestartSnoozeScheduleInDays != *input.EngagedRestartSnoozeScheduleInDays) {
		return false
	}

	if p.EngagedRestartTransitionScheduleInDays != nil && (input.EngagedRestartTransitionScheduleInDays == nil || *p.EngagedRestartTransitionScheduleInDays != *input.EngagedRestartTransitionScheduleInDays) {
		return false
	}

	if p.FeatureUpdatesDeferralPeriodInDays != nil && (input.FeatureUpdatesDeferralPeriodInDays == nil || *p.FeatureUpdatesDeferralPeriodInDays != *input.FeatureUpdatesDeferralPeriodInDays) {
		return false
	}

	if p.FeatureUpdatesPauseExpiryDateTime != nil && (input.FeatureUpdatesPauseExpiryDateTime == nil || *p.FeatureUpdatesPauseExpiryDateTime != *input.FeatureUpdatesPauseExpiryDateTime) {
		return false
	}

	if p.FeatureUpdatesPauseStartDate != nil && (input.FeatureUpdatesPauseStartDate == nil || *p.FeatureUpdatesPauseStartDate != *input.FeatureUpdatesPauseStartDate) {
		return false
	}

	if p.FeatureUpdatesPaused != nil && (input.FeatureUpdatesPaused == nil || *p.FeatureUpdatesPaused != *input.FeatureUpdatesPaused) {
		return false
	}

	if p.FeatureUpdatesRollbackStartDateTime != nil && (input.FeatureUpdatesRollbackStartDateTime == nil || *p.FeatureUpdatesRollbackStartDateTime != *input.FeatureUpdatesRollbackStartDateTime) {
		return false
	}

	if p.FeatureUpdatesRollbackWindowInDays != nil && (input.FeatureUpdatesRollbackWindowInDays == nil || *p.FeatureUpdatesRollbackWindowInDays != *input.FeatureUpdatesRollbackWindowInDays) {
		return false
	}

	if p.FeatureUpdatesWillBeRolledBack != nil && (input.FeatureUpdatesWillBeRolledBack == nil || *p.FeatureUpdatesWillBeRolledBack != *input.FeatureUpdatesWillBeRolledBack) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MicrosoftUpdateServiceAllowed != nil && (input.MicrosoftUpdateServiceAllowed == nil || *p.MicrosoftUpdateServiceAllowed != *input.MicrosoftUpdateServiceAllowed) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PostponeRebootUntilAfterDeadline != nil && (input.PostponeRebootUntilAfterDeadline == nil || *p.PostponeRebootUntilAfterDeadline != *input.PostponeRebootUntilAfterDeadline) {
		return false
	}

	if p.QualityUpdatesDeferralPeriodInDays != nil && (input.QualityUpdatesDeferralPeriodInDays == nil || *p.QualityUpdatesDeferralPeriodInDays != *input.QualityUpdatesDeferralPeriodInDays) {
		return false
	}

	if p.QualityUpdatesPauseExpiryDateTime != nil && (input.QualityUpdatesPauseExpiryDateTime == nil || *p.QualityUpdatesPauseExpiryDateTime != *input.QualityUpdatesPauseExpiryDateTime) {
		return false
	}

	if p.QualityUpdatesPauseStartDate != nil && (input.QualityUpdatesPauseStartDate == nil || *p.QualityUpdatesPauseStartDate != *input.QualityUpdatesPauseStartDate) {
		return false
	}

	if p.QualityUpdatesPaused != nil && (input.QualityUpdatesPaused == nil || *p.QualityUpdatesPaused != *input.QualityUpdatesPaused) {
		return false
	}

	if p.QualityUpdatesRollbackStartDateTime != nil && (input.QualityUpdatesRollbackStartDateTime == nil || *p.QualityUpdatesRollbackStartDateTime != *input.QualityUpdatesRollbackStartDateTime) {
		return false
	}

	if p.QualityUpdatesWillBeRolledBack != nil && (input.QualityUpdatesWillBeRolledBack == nil || *p.QualityUpdatesWillBeRolledBack != *input.QualityUpdatesWillBeRolledBack) {
		return false
	}

	if p.ScheduleImminentRestartWarningInMinutes != nil && (input.ScheduleImminentRestartWarningInMinutes == nil || *p.ScheduleImminentRestartWarningInMinutes != *input.ScheduleImminentRestartWarningInMinutes) {
		return false
	}

	if p.ScheduleRestartWarningInHours != nil && (input.ScheduleRestartWarningInHours == nil || *p.ScheduleRestartWarningInHours != *input.ScheduleRestartWarningInHours) {
		return false
	}

	if p.SkipChecksBeforeRestart != nil && (input.SkipChecksBeforeRestart == nil || *p.SkipChecksBeforeRestart != *input.SkipChecksBeforeRestart) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
