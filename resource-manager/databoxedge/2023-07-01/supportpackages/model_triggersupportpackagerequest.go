package supportpackages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerSupportPackageRequest struct {
	Id         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties SupportPackageRequestProperties `json:"properties"`
	Type       *string                         `json:"type,omitempty"`
}
