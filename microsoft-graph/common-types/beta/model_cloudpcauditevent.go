package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEvent struct {
	Activity              *string                                 `json:"activity,omitempty"`
	ActivityDateTime      *string                                 `json:"activityDateTime,omitempty"`
	ActivityOperationType *CloudPCAuditEventActivityOperationType `json:"activityOperationType,omitempty"`
	ActivityResult        *CloudPCAuditEventActivityResult        `json:"activityResult,omitempty"`
	ActivityType          *string                                 `json:"activityType,omitempty"`
	Actor                 *CloudPCAuditActor                      `json:"actor,omitempty"`
	Category              *CloudPCAuditEventCategory              `json:"category,omitempty"`
	ComponentName         *string                                 `json:"componentName,omitempty"`
	CorrelationId         *string                                 `json:"correlationId,omitempty"`
	DisplayName           *string                                 `json:"displayName,omitempty"`
	Id                    *string                                 `json:"id,omitempty"`
	ODataType             *string                                 `json:"@odata.type,omitempty"`
	Resources             *[]CloudPCAuditResource                 `json:"resources,omitempty"`
}
