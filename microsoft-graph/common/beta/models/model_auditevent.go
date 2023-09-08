package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditEvent struct {
	Activity              *string          `json:"activity,omitempty"`
	ActivityDateTime      *string          `json:"activityDateTime,omitempty"`
	ActivityOperationType *string          `json:"activityOperationType,omitempty"`
	ActivityResult        *string          `json:"activityResult,omitempty"`
	ActivityType          *string          `json:"activityType,omitempty"`
	Actor                 *AuditActor      `json:"actor,omitempty"`
	Category              *string          `json:"category,omitempty"`
	ComponentName         *string          `json:"componentName,omitempty"`
	CorrelationId         *string          `json:"correlationId,omitempty"`
	DisplayName           *string          `json:"displayName,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	ODataType             *string          `json:"@odata.type,omitempty"`
	Resources             *[]AuditResource `json:"resources,omitempty"`
}
