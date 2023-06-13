package recoverypointscrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointMoveReadinessInfo struct {
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	IsReadyForMove *bool   `json:"isReadyForMove,omitempty"`
}
