package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitlockerRecoveryKey struct {
	CreatedDateTime *string                         `json:"createdDateTime,omitempty"`
	DeviceId        *string                         `json:"deviceId,omitempty"`
	Id              *string                         `json:"id,omitempty"`
	Key             *string                         `json:"key,omitempty"`
	ODataType       *string                         `json:"@odata.type,omitempty"`
	VolumeType      *BitlockerRecoveryKeyVolumeType `json:"volumeType,omitempty"`
}
