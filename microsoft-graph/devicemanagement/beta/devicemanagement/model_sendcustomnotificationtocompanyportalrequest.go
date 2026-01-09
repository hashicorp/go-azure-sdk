package devicemanagement

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendCustomNotificationToCompanyPortalRequest struct {
	GroupsToNotify    *[]string             `json:"groupsToNotify,omitempty"`
	NotificationBody  nullable.Type[string] `json:"notificationBody,omitempty"`
	NotificationTitle nullable.Type[string] `json:"notificationTitle,omitempty"`
}
