package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentRequestPolicy struct {
	Id                    *string                      `json:"id,omitempty"`
	IsEnabled             *bool                        `json:"isEnabled,omitempty"`
	NotifyReviewers       *bool                        `json:"notifyReviewers,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	RemindersEnabled      *bool                        `json:"remindersEnabled,omitempty"`
	RequestDurationInDays *int64                       `json:"requestDurationInDays,omitempty"`
	Reviewers             *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`
	Version               *int64                       `json:"version,omitempty"`
}
