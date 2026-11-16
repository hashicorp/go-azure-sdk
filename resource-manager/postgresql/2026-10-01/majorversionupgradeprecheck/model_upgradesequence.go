package majorversionupgradeprecheck

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpgradeSequence struct {
	SourceVersion *string `json:"sourceVersion,omitempty"`
	TargetVersion *string `json:"targetVersion,omitempty"`
}
