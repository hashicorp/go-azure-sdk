package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNetworkUsageRule struct {
	CellularDataBlockWhenRoaming *bool          `json:"cellularDataBlockWhenRoaming,omitempty"`
	CellularDataBlocked          *bool          `json:"cellularDataBlocked,omitempty"`
	ManagedApps                  *[]AppListItem `json:"managedApps,omitempty"`
	ODataType                    *string        `json:"@odata.type,omitempty"`
}
