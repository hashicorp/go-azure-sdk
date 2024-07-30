package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSetting struct {
	ExternalAudience       *AutomaticRepliesSettingExternalAudience `json:"externalAudience,omitempty"`
	ExternalReplyMessage   *string                                  `json:"externalReplyMessage,omitempty"`
	InternalReplyMessage   *string                                  `json:"internalReplyMessage,omitempty"`
	ODataType              *string                                  `json:"@odata.type,omitempty"`
	ScheduledEndDateTime   *DateTimeTimeZone                        `json:"scheduledEndDateTime,omitempty"`
	ScheduledStartDateTime *DateTimeTimeZone                        `json:"scheduledStartDateTime,omitempty"`
	Status                 *AutomaticRepliesSettingStatus           `json:"status,omitempty"`
}
