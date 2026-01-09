package driveitem

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyDriveItemRequest struct {
	ChildrenOnly             nullable.Type[bool]   `json:"childrenOnly,omitempty"`
	IncludeAllVersionHistory nullable.Type[bool]   `json:"includeAllVersionHistory,omitempty"`
	Name                     nullable.Type[string] `json:"name,omitempty"`
	ParentReference          *beta.ItemReference   `json:"parentReference,omitempty"`
}
