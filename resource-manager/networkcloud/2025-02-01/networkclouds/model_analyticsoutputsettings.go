package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticsOutputSettings struct {
	AnalyticsWorkspaceId *string           `json:"analyticsWorkspaceId,omitempty"`
	AssociatedIdentity   *IdentitySelector `json:"associatedIdentity,omitempty"`
}
