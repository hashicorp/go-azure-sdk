package devicemanagement

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendCustomNotificationToCompanyPortalRequest struct {
	GroupsToNotify    *[]string             `json:"groupsToNotify,omitempty"`
	NotificationBody  nullable.Type[string] `json:"notificationBody,omitempty"`
	NotificationTitle nullable.Type[string] `json:"notificationTitle,omitempty"`
}
