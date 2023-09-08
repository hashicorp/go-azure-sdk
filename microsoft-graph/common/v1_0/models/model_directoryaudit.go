package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAudit struct {
	ActivityDateTime    *string                 `json:"activityDateTime,omitempty"`
	ActivityDisplayName *string                 `json:"activityDisplayName,omitempty"`
	AdditionalDetails   *[]KeyValue             `json:"additionalDetails,omitempty"`
	Category            *string                 `json:"category,omitempty"`
	CorrelationId       *string                 `json:"correlationId,omitempty"`
	Id                  *string                 `json:"id,omitempty"`
	InitiatedBy         *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	LoggedByService     *string                 `json:"loggedByService,omitempty"`
	ODataType           *string                 `json:"@odata.type,omitempty"`
	OperationType       *string                 `json:"operationType,omitempty"`
	Result              *DirectoryAuditResult   `json:"result,omitempty"`
	ResultReason        *string                 `json:"resultReason,omitempty"`
	TargetResources     *[]TargetResource       `json:"targetResources,omitempty"`
}
