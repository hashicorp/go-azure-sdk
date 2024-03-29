package collection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAcrConfigurationInfo struct {
	LoginServers *[]string                  `json:"loginServers,omitempty"`
	OciArtifacts *[]ServiceOciArtifactEntry `json:"ociArtifacts,omitempty"`
}
