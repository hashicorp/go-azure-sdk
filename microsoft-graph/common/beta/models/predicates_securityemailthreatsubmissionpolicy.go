package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionPolicyOperationPredicate struct {
	CustomizedNotificationSenderEmailAddress *string
	CustomizedReportRecipientEmailAddress    *string
	Id                                       *string
	IsAlwaysReportEnabledForUsers            *bool
	IsAskMeEnabledForUsers                   *bool
	IsCustomizedMessageEnabled               *bool
	IsCustomizedMessageEnabledForPhishing    *bool
	IsCustomizedNotificationSenderEnabled    *bool
	IsNeverReportEnabledForUsers             *bool
	IsOrganizationBrandingEnabled            *bool
	IsReportFromQuarantineEnabled            *bool
	IsReportToCustomizedEmailAddressEnabled  *bool
	IsReportToMicrosoftEnabled               *bool
	IsReviewEmailNotificationEnabled         *bool
	ODataType                                *string
}

func (p SecurityEmailThreatSubmissionPolicyOperationPredicate) Matches(input SecurityEmailThreatSubmissionPolicy) bool {

	if p.CustomizedNotificationSenderEmailAddress != nil && (input.CustomizedNotificationSenderEmailAddress == nil || *p.CustomizedNotificationSenderEmailAddress != *input.CustomizedNotificationSenderEmailAddress) {
		return false
	}

	if p.CustomizedReportRecipientEmailAddress != nil && (input.CustomizedReportRecipientEmailAddress == nil || *p.CustomizedReportRecipientEmailAddress != *input.CustomizedReportRecipientEmailAddress) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAlwaysReportEnabledForUsers != nil && (input.IsAlwaysReportEnabledForUsers == nil || *p.IsAlwaysReportEnabledForUsers != *input.IsAlwaysReportEnabledForUsers) {
		return false
	}

	if p.IsAskMeEnabledForUsers != nil && (input.IsAskMeEnabledForUsers == nil || *p.IsAskMeEnabledForUsers != *input.IsAskMeEnabledForUsers) {
		return false
	}

	if p.IsCustomizedMessageEnabled != nil && (input.IsCustomizedMessageEnabled == nil || *p.IsCustomizedMessageEnabled != *input.IsCustomizedMessageEnabled) {
		return false
	}

	if p.IsCustomizedMessageEnabledForPhishing != nil && (input.IsCustomizedMessageEnabledForPhishing == nil || *p.IsCustomizedMessageEnabledForPhishing != *input.IsCustomizedMessageEnabledForPhishing) {
		return false
	}

	if p.IsCustomizedNotificationSenderEnabled != nil && (input.IsCustomizedNotificationSenderEnabled == nil || *p.IsCustomizedNotificationSenderEnabled != *input.IsCustomizedNotificationSenderEnabled) {
		return false
	}

	if p.IsNeverReportEnabledForUsers != nil && (input.IsNeverReportEnabledForUsers == nil || *p.IsNeverReportEnabledForUsers != *input.IsNeverReportEnabledForUsers) {
		return false
	}

	if p.IsOrganizationBrandingEnabled != nil && (input.IsOrganizationBrandingEnabled == nil || *p.IsOrganizationBrandingEnabled != *input.IsOrganizationBrandingEnabled) {
		return false
	}

	if p.IsReportFromQuarantineEnabled != nil && (input.IsReportFromQuarantineEnabled == nil || *p.IsReportFromQuarantineEnabled != *input.IsReportFromQuarantineEnabled) {
		return false
	}

	if p.IsReportToCustomizedEmailAddressEnabled != nil && (input.IsReportToCustomizedEmailAddressEnabled == nil || *p.IsReportToCustomizedEmailAddressEnabled != *input.IsReportToCustomizedEmailAddressEnabled) {
		return false
	}

	if p.IsReportToMicrosoftEnabled != nil && (input.IsReportToMicrosoftEnabled == nil || *p.IsReportToMicrosoftEnabled != *input.IsReportToMicrosoftEnabled) {
		return false
	}

	if p.IsReviewEmailNotificationEnabled != nil && (input.IsReviewEmailNotificationEnabled == nil || *p.IsReviewEmailNotificationEnabled != *input.IsReviewEmailNotificationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
