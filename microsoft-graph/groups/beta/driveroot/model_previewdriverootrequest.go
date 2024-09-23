package driveroot

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreviewDriveRootRequest struct {
	AllowEdit  nullable.Type[bool]   `json:"allowEdit,omitempty"`
	Chromeless nullable.Type[bool]   `json:"chromeless,omitempty"`
	Page       nullable.Type[string] `json:"page,omitempty"`
	Viewer     nullable.Type[string] `json:"viewer,omitempty"`
}
