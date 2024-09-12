package comanageddevice

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComanagedDeviceBulkRestoreCloudPCRequest struct {
	ManagedDeviceIds     *[]string              `json:"managedDeviceIds,omitempty"`
	RestorePointDateTime nullable.Type[string]  `json:"restorePointDateTime,omitempty"`
	TimeRange            *beta.RestoreTimeRange `json:"timeRange,omitempty"`
}
