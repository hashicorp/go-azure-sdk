package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentProperties struct {
	Active      *bool   `json:"active,omitempty"`
	Author      *string `json:"author,omitempty"`
	AuthorEmail *string `json:"author_email,omitempty"`
	Deployer    *string `json:"deployer,omitempty"`
	Details     *string `json:"details,omitempty"`
	EndTime     *string `json:"end_time,omitempty"`
	Message     *string `json:"message,omitempty"`
	StartTime   *string `json:"start_time,omitempty"`
	Status      *int64  `json:"status,omitempty"`
}
