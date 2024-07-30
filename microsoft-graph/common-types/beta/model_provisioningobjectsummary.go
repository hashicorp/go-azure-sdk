package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningObjectSummary struct {
	Action                 *string                                      `json:"action,omitempty"`
	ActivityDateTime       *string                                      `json:"activityDateTime,omitempty"`
	ChangeId               *string                                      `json:"changeId,omitempty"`
	CycleId                *string                                      `json:"cycleId,omitempty"`
	DurationInMilliseconds *int64                                       `json:"durationInMilliseconds,omitempty"`
	Id                     *string                                      `json:"id,omitempty"`
	InitiatedBy            *Initiator                                   `json:"initiatedBy,omitempty"`
	JobId                  *string                                      `json:"jobId,omitempty"`
	ModifiedProperties     *[]ModifiedProperty                          `json:"modifiedProperties,omitempty"`
	ODataType              *string                                      `json:"@odata.type,omitempty"`
	ProvisioningAction     *ProvisioningObjectSummaryProvisioningAction `json:"provisioningAction,omitempty"`
	ProvisioningStatusInfo *ProvisioningStatusInfo                      `json:"provisioningStatusInfo,omitempty"`
	ProvisioningSteps      *[]ProvisioningStep                          `json:"provisioningSteps,omitempty"`
	ServicePrincipal       *ProvisioningServicePrincipal                `json:"servicePrincipal,omitempty"`
	SourceIdentity         *ProvisionedIdentity                         `json:"sourceIdentity,omitempty"`
	SourceSystem           *ProvisioningSystem                          `json:"sourceSystem,omitempty"`
	StatusInfo             *StatusBase                                  `json:"statusInfo,omitempty"`
	TargetIdentity         *ProvisionedIdentity                         `json:"targetIdentity,omitempty"`
	TargetSystem           *ProvisioningSystem                          `json:"targetSystem,omitempty"`
	TenantId               *string                                      `json:"tenantId,omitempty"`
}
