package notificationrecipientuser

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationName string

const (
	NotificationNameAccountClosedPublisher                            NotificationName = "AccountClosedPublisher"
	NotificationNameBCC                                               NotificationName = "BCC"
	NotificationNameNewApplicationNotificationMessage                 NotificationName = "NewApplicationNotificationMessage"
	NotificationNameNewIssuePublisherNotificationMessage              NotificationName = "NewIssuePublisherNotificationMessage"
	NotificationNamePurchasePublisherNotificationMessage              NotificationName = "PurchasePublisherNotificationMessage"
	NotificationNameQuotaLimitApproachingPublisherNotificationMessage NotificationName = "QuotaLimitApproachingPublisherNotificationMessage"
	NotificationNameRequestPublisherNotificationMessage               NotificationName = "RequestPublisherNotificationMessage"
)

func PossibleValuesForNotificationName() []string {
	return []string{
		string(NotificationNameAccountClosedPublisher),
		string(NotificationNameBCC),
		string(NotificationNameNewApplicationNotificationMessage),
		string(NotificationNameNewIssuePublisherNotificationMessage),
		string(NotificationNamePurchasePublisherNotificationMessage),
		string(NotificationNameQuotaLimitApproachingPublisherNotificationMessage),
		string(NotificationNameRequestPublisherNotificationMessage),
	}
}
