package tenantattachrbac

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableTenantAttachRBACRequest struct {
	Enable *bool `json:"enable,omitempty"`
}
