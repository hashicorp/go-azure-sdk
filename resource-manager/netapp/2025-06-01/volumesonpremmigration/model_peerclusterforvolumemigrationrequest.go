package volumesonpremmigration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeerClusterForVolumeMigrationRequest struct {
	PeerIPAddresses []string `json:"peerIpAddresses"`
}
