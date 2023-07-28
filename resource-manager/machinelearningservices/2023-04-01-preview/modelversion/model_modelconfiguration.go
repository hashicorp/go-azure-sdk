package modelversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelConfiguration struct {
	Mode      *PackageInputDeliveryMode `json:"mode,omitempty"`
	MountPath *string                   `json:"mountPath,omitempty"`
}
