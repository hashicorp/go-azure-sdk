package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppAssignmentSettings struct {
	AndroidManagedStoreAppTrackIds *[]string                                               `json:"androidManagedStoreAppTrackIds,omitempty"`
	AutoUpdateMode                 *AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode `json:"autoUpdateMode,omitempty"`
	ODataType                      *string                                                 `json:"@odata.type,omitempty"`
}
