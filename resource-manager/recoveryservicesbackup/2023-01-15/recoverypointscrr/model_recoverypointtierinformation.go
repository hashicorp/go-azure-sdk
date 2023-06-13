package recoverypointscrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointTierInformation struct {
	ExtendedInfo *map[string]string       `json:"extendedInfo,omitempty"`
	Status       *RecoveryPointTierStatus `json:"status,omitempty"`
	Type         *RecoveryPointTierType   `json:"type,omitempty"`
}
