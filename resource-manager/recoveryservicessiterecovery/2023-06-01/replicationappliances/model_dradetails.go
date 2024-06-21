package replicationappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DraDetails struct {
	BiosId                    *string           `json:"biosId,omitempty"`
	ForwardProtectedItemCount *int64            `json:"forwardProtectedItemCount,omitempty"`
	Health                    *ProtectionHealth `json:"health,omitempty"`
	HealthErrors              *[]HealthError    `json:"healthErrors,omitempty"`
	Id                        *string           `json:"id,omitempty"`
	LastHeartbeatUtc          *string           `json:"lastHeartbeatUtc,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	ReverseProtectedItemCount *int64            `json:"reverseProtectedItemCount,omitempty"`
	Version                   *string           `json:"version,omitempty"`
}
