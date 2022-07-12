package disks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageDiskReference struct {
	Id  string `json:"id"`
	Lun *int64 `json:"lun,omitempty"`
}
