package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlert struct {
	AlertData                 *ManagedTenantsAlertData                        `json:"alertData,omitempty"`
	AlertDataReferenceStrings *[]ManagedTenantsAlertDataReferenceString       `json:"alertDataReferenceStrings,omitempty"`
	AlertLogs                 *[]ManagedTenantsManagedTenantAlertLog          `json:"alertLogs,omitempty"`
	AlertRule                 *ManagedTenantsManagedTenantAlertRule           `json:"alertRule,omitempty"`
	AlertRuleDisplayName      *string                                         `json:"alertRuleDisplayName,omitempty"`
	ApiNotifications          *[]ManagedTenantsManagedTenantApiNotification   `json:"apiNotifications,omitempty"`
	AssignedToUserId          *string                                         `json:"assignedToUserId,omitempty"`
	CorrelationCount          *int64                                          `json:"correlationCount,omitempty"`
	CorrelationId             *string                                         `json:"correlationId,omitempty"`
	CreatedByUserId           *string                                         `json:"createdByUserId,omitempty"`
	CreatedDateTime           *string                                         `json:"createdDateTime,omitempty"`
	EmailNotifications        *[]ManagedTenantsManagedTenantEmailNotification `json:"emailNotifications,omitempty"`
	Id                        *string                                         `json:"id,omitempty"`
	LastActionByUserId        *string                                         `json:"lastActionByUserId,omitempty"`
	LastActionDateTime        *string                                         `json:"lastActionDateTime,omitempty"`
	Message                   *string                                         `json:"message,omitempty"`
	ODataType                 *string                                         `json:"@odata.type,omitempty"`
	Severity                  *ManagedTenantsManagedTenantAlertSeverity       `json:"severity,omitempty"`
	Status                    *ManagedTenantsManagedTenantAlertStatus         `json:"status,omitempty"`
	TenantId                  *string                                         `json:"tenantId,omitempty"`
	Title                     *string                                         `json:"title,omitempty"`
}
