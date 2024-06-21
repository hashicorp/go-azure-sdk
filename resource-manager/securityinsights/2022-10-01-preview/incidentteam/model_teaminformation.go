package incidentteam

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamInformation struct {
	Description         *string `json:"description,omitempty"`
	Name                *string `json:"name,omitempty"`
	PrimaryChannelUrl   *string `json:"primaryChannelUrl,omitempty"`
	TeamCreationTimeUtc *string `json:"teamCreationTimeUtc,omitempty"`
	TeamId              *string `json:"teamId,omitempty"`
}
