package volumegroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumeGroupUpdateProperties struct {
	Encryption   *EncryptionType    `json:"encryption,omitempty"`
	NetworkAcls  *NetworkRuleSet    `json:"networkAcls"`
	ProtocolType *StorageTargetType `json:"protocolType,omitempty"`
}
