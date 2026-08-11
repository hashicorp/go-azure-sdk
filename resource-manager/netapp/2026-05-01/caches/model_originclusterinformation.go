package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OriginClusterInformation struct {
	PeerAddresses   []string `json:"peerAddresses"`
	PeerClusterName string   `json:"peerClusterName"`
	PeerVolumeName  string   `json:"peerVolumeName"`
	PeerVserverName string   `json:"peerVserverName"`
}
