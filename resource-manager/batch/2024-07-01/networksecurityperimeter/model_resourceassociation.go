package networksecurityperimeter

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAssociation struct {
	AccessMode *ResourceAssociationAccessMode `json:"accessMode,omitempty"`
	Name       *string                        `json:"name,omitempty"`
}
