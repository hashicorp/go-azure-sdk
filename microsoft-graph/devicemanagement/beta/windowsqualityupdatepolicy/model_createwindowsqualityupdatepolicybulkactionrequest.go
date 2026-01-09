package windowsqualityupdatepolicy

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsQualityUpdatePolicyBulkActionRequest struct {
	// An enum type to represent approval actions of single or list of quality update policies
	Action *beta.WindowsQualityUpdatePolicyActionType `json:"action,omitempty"`

	Ids *[]string `json:"ids,omitempty"`
}
