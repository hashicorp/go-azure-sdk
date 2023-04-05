package portalrevision

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalRevisionStatus string

const (
	PortalRevisionStatusCompleted  PortalRevisionStatus = "completed"
	PortalRevisionStatusFailed     PortalRevisionStatus = "failed"
	PortalRevisionStatusPending    PortalRevisionStatus = "pending"
	PortalRevisionStatusPublishing PortalRevisionStatus = "publishing"
)

func PossibleValuesForPortalRevisionStatus() []string {
	return []string{
		string(PortalRevisionStatusCompleted),
		string(PortalRevisionStatusFailed),
		string(PortalRevisionStatusPending),
		string(PortalRevisionStatusPublishing),
	}
}
