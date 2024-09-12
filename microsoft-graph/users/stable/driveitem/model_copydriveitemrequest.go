package driveitem

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyDriveItemRequest struct {
	Name            nullable.Type[string] `json:"name,omitempty"`
	ParentReference *stable.ItemReference `json:"parentReference,omitempty"`
}
