package emailtemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateName string

const (
	TemplateNameAccountClosedDeveloper                            TemplateName = "accountClosedDeveloper"
	TemplateNameApplicationApprovedNotificationMessage            TemplateName = "applicationApprovedNotificationMessage"
	TemplateNameConfirmSignUpIdentityDefault                      TemplateName = "confirmSignUpIdentityDefault"
	TemplateNameEmailChangeIdentityDefault                        TemplateName = "emailChangeIdentityDefault"
	TemplateNameInviteUserNotificationMessage                     TemplateName = "inviteUserNotificationMessage"
	TemplateNameNewCommentNotificationMessage                     TemplateName = "newCommentNotificationMessage"
	TemplateNameNewDeveloperNotificationMessage                   TemplateName = "newDeveloperNotificationMessage"
	TemplateNameNewIssueNotificationMessage                       TemplateName = "newIssueNotificationMessage"
	TemplateNamePasswordResetByAdminNotificationMessage           TemplateName = "passwordResetByAdminNotificationMessage"
	TemplateNamePasswordResetIdentityDefault                      TemplateName = "passwordResetIdentityDefault"
	TemplateNamePurchaseDeveloperNotificationMessage              TemplateName = "purchaseDeveloperNotificationMessage"
	TemplateNameQuotaLimitApproachingDeveloperNotificationMessage TemplateName = "quotaLimitApproachingDeveloperNotificationMessage"
	TemplateNameRejectDeveloperNotificationMessage                TemplateName = "rejectDeveloperNotificationMessage"
	TemplateNameRequestDeveloperNotificationMessage               TemplateName = "requestDeveloperNotificationMessage"
)

func PossibleValuesForTemplateName() []string {
	return []string{
		string(TemplateNameAccountClosedDeveloper),
		string(TemplateNameApplicationApprovedNotificationMessage),
		string(TemplateNameConfirmSignUpIdentityDefault),
		string(TemplateNameEmailChangeIdentityDefault),
		string(TemplateNameInviteUserNotificationMessage),
		string(TemplateNameNewCommentNotificationMessage),
		string(TemplateNameNewDeveloperNotificationMessage),
		string(TemplateNameNewIssueNotificationMessage),
		string(TemplateNamePasswordResetByAdminNotificationMessage),
		string(TemplateNamePasswordResetIdentityDefault),
		string(TemplateNamePurchaseDeveloperNotificationMessage),
		string(TemplateNameQuotaLimitApproachingDeveloperNotificationMessage),
		string(TemplateNameRejectDeveloperNotificationMessage),
		string(TemplateNameRequestDeveloperNotificationMessage),
	}
}
