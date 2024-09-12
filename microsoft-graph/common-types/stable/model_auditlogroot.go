package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuditLogRoot{}

type AuditLogRoot struct {
	DirectoryAudits *[]DirectoryAudit            `json:"directoryAudits,omitempty"`
	Provisioning    *[]ProvisioningObjectSummary `json:"provisioning,omitempty"`
	SignIns         *[]SignIn                    `json:"signIns,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AuditLogRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuditLogRoot{}

func (s AuditLogRoot) MarshalJSON() ([]byte, error) {
	type wrapper AuditLogRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuditLogRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuditLogRoot: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.auditLogRoot"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuditLogRoot: %+v", err)
	}

	return encoded, nil
}
