package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCAuditEvent{}

type CloudPCAuditEvent struct {
	// The friendly name of the audit activity.
	Activity nullable.Type[string] `json:"activity,omitempty"`

	// The date time in UTC when the activity was performed. Read-only.
	ActivityDateTime *string `json:"activityDateTime,omitempty"`

	ActivityOperationType *CloudPCAuditActivityOperationType `json:"activityOperationType,omitempty"`
	ActivityResult        *CloudPCAuditActivityResult        `json:"activityResult,omitempty"`

	// The type of activity that was performed. Read-only.
	ActivityType *string `json:"activityType,omitempty"`

	Actor    *CloudPCAuditActor    `json:"actor,omitempty"`
	Category *CloudPCAuditCategory `json:"category,omitempty"`

	// The component name for the audit event. Read-only.
	ComponentName *string `json:"componentName,omitempty"`

	// The client request ID that is used to correlate activity within the system. Read-only.
	CorrelationId *string `json:"correlationId,omitempty"`

	// The display name for the audit event. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of cloudPcAuditResource objects. Read-only.
	Resources *[]CloudPCAuditResource `json:"resources,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s CloudPCAuditEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCAuditEvent{}

func (s CloudPCAuditEvent) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCAuditEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCAuditEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCAuditEvent: %+v", err)
	}

	delete(decoded, "componentName")
	delete(decoded, "displayName")
	decoded["@odata.type"] = "#microsoft.graph.cloudPcAuditEvent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCAuditEvent: %+v", err)
	}

	return encoded, nil
}
