package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlag struct {
	CompletedDateTime *DateTimeTimeZone       `json:"completedDateTime,omitempty"`
	DueDateTime       *DateTimeTimeZone       `json:"dueDateTime,omitempty"`
	FlagStatus        *FollowupFlagFlagStatus `json:"flagStatus,omitempty"`
	ODataType         *string                 `json:"@odata.type,omitempty"`
	StartDateTime     *DateTimeTimeZone       `json:"startDateTime,omitempty"`
}
