package recoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskInformation struct {
	Lun  *int64  `json:"lun,omitempty"`
	Name *string `json:"name,omitempty"`
}
