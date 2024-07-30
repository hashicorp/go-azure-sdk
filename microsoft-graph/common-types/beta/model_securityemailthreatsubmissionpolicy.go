package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionPolicy struct {
	CustomizedNotificationSenderEmailAddress *string `json:"customizedNotificationSenderEmailAddress,omitempty"`
	CustomizedReportRecipientEmailAddress    *string `json:"customizedReportRecipientEmailAddress,omitempty"`
	Id                                       *string `json:"id,omitempty"`
	IsAlwaysReportEnabledForUsers            *bool   `json:"isAlwaysReportEnabledForUsers,omitempty"`
	IsAskMeEnabledForUsers                   *bool   `json:"isAskMeEnabledForUsers,omitempty"`
	IsCustomizedMessageEnabled               *bool   `json:"isCustomizedMessageEnabled,omitempty"`
	IsCustomizedMessageEnabledForPhishing    *bool   `json:"isCustomizedMessageEnabledForPhishing,omitempty"`
	IsCustomizedNotificationSenderEnabled    *bool   `json:"isCustomizedNotificationSenderEnabled,omitempty"`
	IsNeverReportEnabledForUsers             *bool   `json:"isNeverReportEnabledForUsers,omitempty"`
	IsOrganizationBrandingEnabled            *bool   `json:"isOrganizationBrandingEnabled,omitempty"`
	IsReportFromQuarantineEnabled            *bool   `json:"isReportFromQuarantineEnabled,omitempty"`
	IsReportToCustomizedEmailAddressEnabled  *bool   `json:"isReportToCustomizedEmailAddressEnabled,omitempty"`
	IsReportToMicrosoftEnabled               *bool   `json:"isReportToMicrosoftEnabled,omitempty"`
	IsReviewEmailNotificationEnabled         *bool   `json:"isReviewEmailNotificationEnabled,omitempty"`
	ODataType                                *string `json:"@odata.type,omitempty"`
}
