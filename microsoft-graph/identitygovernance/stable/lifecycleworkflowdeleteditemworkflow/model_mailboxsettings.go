package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettings struct {
	ArchiveFolder                         *string                                               `json:"archiveFolder,omitempty"`
	AutomaticRepliesSetting               *AutomaticRepliesSetting                              `json:"automaticRepliesSetting,omitempty"`
	DateFormat                            *string                                               `json:"dateFormat,omitempty"`
	DelegateMeetingMessageDeliveryOptions *MailboxSettingsDelegateMeetingMessageDeliveryOptions `json:"delegateMeetingMessageDeliveryOptions,omitempty"`
	Language                              *LocaleInfo                                           `json:"language,omitempty"`
	ODataType                             *string                                               `json:"@odata.type,omitempty"`
	TimeFormat                            *string                                               `json:"timeFormat,omitempty"`
	TimeZone                              *string                                               `json:"timeZone,omitempty"`
	UserPurpose                           *MailboxSettingsUserPurpose                           `json:"userPurpose,omitempty"`
	WorkingHours                          *WorkingHours                                         `json:"workingHours,omitempty"`
}
