package recoverypoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointDataStoreDetails struct {
	CreationTime          *string            `json:"creationTime,omitempty"`
	ExpiryTime            *string            `json:"expiryTime,omitempty"`
	Id                    *string            `json:"id,omitempty"`
	MetaData              *string            `json:"metaData,omitempty"`
	RehydrationExpiryTime *string            `json:"rehydrationExpiryTime,omitempty"`
	RehydrationStatus     *RehydrationStatus `json:"rehydrationStatus,omitempty"`
	State                 *string            `json:"state,omitempty"`
	Type                  *string            `json:"type,omitempty"`
	Visible               *bool              `json:"visible,omitempty"`
}
