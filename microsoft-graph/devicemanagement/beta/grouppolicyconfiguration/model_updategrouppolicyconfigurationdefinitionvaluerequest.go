package grouppolicyconfiguration

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyConfigurationDefinitionValueRequest struct {
	Added      *[]beta.GroupPolicyDefinitionValue `json:"added,omitempty"`
	DeletedIds *[]string                          `json:"deletedIds,omitempty"`
	Updated    *[]beta.GroupPolicyDefinitionValue `json:"updated,omitempty"`
}
