package storagesyncservicesresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncServiceUpdateProperties struct {
	IncomingTrafficPolicy *IncomingTrafficPolicy `json:"incomingTrafficPolicy,omitempty"`
	UseIdentity           *bool                  `json:"useIdentity,omitempty"`
}
