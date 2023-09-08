package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventOperationPredicate struct {
	AllowNewTimeProposals      *bool
	BodyPreview                *string
	ChangeKey                  *string
	CreatedDateTime            *string
	HasAttachments             *bool
	HideAttendees              *bool
	Id                         *string
	IsAllDay                   *bool
	IsCancelled                *bool
	IsDraft                    *bool
	IsOnlineMeeting            *bool
	IsOrganizer                *bool
	IsReminderOn               *bool
	LastModifiedDateTime       *string
	ODataType                  *string
	OccurrenceId               *string
	OnlineMeetingUrl           *string
	OriginalEndTimeZone        *string
	OriginalStart              *string
	OriginalStartTimeZone      *string
	ReminderMinutesBeforeStart *int64
	ResponseRequested          *bool
	SeriesMasterId             *string
	Subject                    *string
	TransactionId              *string
	Uid                        *string
	WebLink                    *string
}

func (p EventOperationPredicate) Matches(input Event) bool {

	if p.AllowNewTimeProposals != nil && (input.AllowNewTimeProposals == nil || *p.AllowNewTimeProposals != *input.AllowNewTimeProposals) {
		return false
	}

	if p.BodyPreview != nil && (input.BodyPreview == nil || *p.BodyPreview != *input.BodyPreview) {
		return false
	}

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.HasAttachments != nil && (input.HasAttachments == nil || *p.HasAttachments != *input.HasAttachments) {
		return false
	}

	if p.HideAttendees != nil && (input.HideAttendees == nil || *p.HideAttendees != *input.HideAttendees) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAllDay != nil && (input.IsAllDay == nil || *p.IsAllDay != *input.IsAllDay) {
		return false
	}

	if p.IsCancelled != nil && (input.IsCancelled == nil || *p.IsCancelled != *input.IsCancelled) {
		return false
	}

	if p.IsDraft != nil && (input.IsDraft == nil || *p.IsDraft != *input.IsDraft) {
		return false
	}

	if p.IsOnlineMeeting != nil && (input.IsOnlineMeeting == nil || *p.IsOnlineMeeting != *input.IsOnlineMeeting) {
		return false
	}

	if p.IsOrganizer != nil && (input.IsOrganizer == nil || *p.IsOrganizer != *input.IsOrganizer) {
		return false
	}

	if p.IsReminderOn != nil && (input.IsReminderOn == nil || *p.IsReminderOn != *input.IsReminderOn) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OccurrenceId != nil && (input.OccurrenceId == nil || *p.OccurrenceId != *input.OccurrenceId) {
		return false
	}

	if p.OnlineMeetingUrl != nil && (input.OnlineMeetingUrl == nil || *p.OnlineMeetingUrl != *input.OnlineMeetingUrl) {
		return false
	}

	if p.OriginalEndTimeZone != nil && (input.OriginalEndTimeZone == nil || *p.OriginalEndTimeZone != *input.OriginalEndTimeZone) {
		return false
	}

	if p.OriginalStart != nil && (input.OriginalStart == nil || *p.OriginalStart != *input.OriginalStart) {
		return false
	}

	if p.OriginalStartTimeZone != nil && (input.OriginalStartTimeZone == nil || *p.OriginalStartTimeZone != *input.OriginalStartTimeZone) {
		return false
	}

	if p.ReminderMinutesBeforeStart != nil && (input.ReminderMinutesBeforeStart == nil || *p.ReminderMinutesBeforeStart != *input.ReminderMinutesBeforeStart) {
		return false
	}

	if p.ResponseRequested != nil && (input.ResponseRequested == nil || *p.ResponseRequested != *input.ResponseRequested) {
		return false
	}

	if p.SeriesMasterId != nil && (input.SeriesMasterId == nil || *p.SeriesMasterId != *input.SeriesMasterId) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	if p.TransactionId != nil && (input.TransactionId == nil || *p.TransactionId != *input.TransactionId) {
		return false
	}

	if p.Uid != nil && (input.Uid == nil || *p.Uid != *input.Uid) {
		return false
	}

	if p.WebLink != nil && (input.WebLink == nil || *p.WebLink != *input.WebLink) {
		return false
	}

	return true
}
