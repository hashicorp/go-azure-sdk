package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertRule struct {
	AlertDisplayName              *string                                                            `json:"alertDisplayName,omitempty"`
	AlertTTL                      *int64                                                             `json:"alertTTL,omitempty"`
	Alerts                        *[]ManagedTenantsManagedTenantAlert                                `json:"alerts,omitempty"`
	CreatedByUserId               *string                                                            `json:"createdByUserId,omitempty"`
	CreatedDateTime               *string                                                            `json:"createdDateTime,omitempty"`
	Description                   *string                                                            `json:"description,omitempty"`
	DisplayName                   *string                                                            `json:"displayName,omitempty"`
	Id                            *string                                                            `json:"id,omitempty"`
	LastActionByUserId            *string                                                            `json:"lastActionByUserId,omitempty"`
	LastActionDateTime            *string                                                            `json:"lastActionDateTime,omitempty"`
	LastRunDateTime               *string                                                            `json:"lastRunDateTime,omitempty"`
	NotificationFinalDestinations *ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations `json:"notificationFinalDestinations,omitempty"`
	ODataType                     *string                                                            `json:"@odata.type,omitempty"`
	RuleDefinition                *ManagedTenantsManagedTenantAlertRuleDefinition                    `json:"ruleDefinition,omitempty"`
	Severity                      *ManagedTenantsManagedTenantAlertRuleSeverity                      `json:"severity,omitempty"`
	Targets                       *[]ManagedTenantsNotificationTarget                                `json:"targets,omitempty"`
	TenantIds                     *[]ManagedTenantsTenantInfo                                        `json:"tenantIds,omitempty"`
}
