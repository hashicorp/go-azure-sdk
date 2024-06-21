package updates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Update struct {
	ImpactDurationInSec *int64            `json:"impactDurationInSec,omitempty"`
	ImpactType          *ImpactType       `json:"impactType,omitempty"`
	MaintenanceScope    *MaintenanceScope `json:"maintenanceScope,omitempty"`
	NotBefore           *string           `json:"notBefore,omitempty"`
	Properties          *UpdateProperties `json:"properties,omitempty"`
	Status              *UpdateStatus     `json:"status,omitempty"`
}
