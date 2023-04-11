package securitycontacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactProperties struct {
	AlertNotifications AlertNotifications `json:"alertNotifications"`
	AlertsToAdmins     AlertsToAdmins     `json:"alertsToAdmins"`
	Email              string             `json:"email"`
	Phone              *string            `json:"phone,omitempty"`
}
