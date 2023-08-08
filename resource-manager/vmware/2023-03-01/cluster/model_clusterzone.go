package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterZone struct {
	Hosts *[]string `json:"hosts,omitempty"`
	Zone  *string   `json:"zone,omitempty"`
}
