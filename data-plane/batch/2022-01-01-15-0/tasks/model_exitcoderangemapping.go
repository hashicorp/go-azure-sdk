package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExitCodeRangeMapping struct {
	End         int64       `json:"end"`
	ExitOptions ExitOptions `json:"exitOptions"`
	Start       int64       `json:"start"`
}
