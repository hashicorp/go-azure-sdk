package notification

import "strings"

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

func parseNotificationName(input string) (*NotificationName, error) {
	vals := map[string]NotificationName{
		"accountclosedpublisher":                            NotificationNameAccountClosedPublisher,
		"bcc":                                               NotificationNameBCC,
		"newapplicationnotificationmessage":                 NotificationNameNewApplicationNotificationMessage,
		"newissuepublishernotificationmessage":              NotificationNameNewIssuePublisherNotificationMessage,
		"purchasepublishernotificationmessage":              NotificationNamePurchasePublisherNotificationMessage,
		"quotalimitapproachingpublishernotificationmessage": NotificationNameQuotaLimitApproachingPublisherNotificationMessage,
		"requestpublishernotificationmessage":               NotificationNameRequestPublisherNotificationMessage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationName(input)
	return &out, nil
}
