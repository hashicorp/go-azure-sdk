package operation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetRestoreInfo struct {
	ContainerId                   *string           `json:"containerId,omitempty"`
	DatabaseName                  *string           `json:"databaseName,omitempty"`
	OverwriteOption               *OverwriteOptions `json:"overwriteOption,omitempty"`
	TargetDirectoryForFileRestore *string           `json:"targetDirectoryForFileRestore,omitempty"`
}
