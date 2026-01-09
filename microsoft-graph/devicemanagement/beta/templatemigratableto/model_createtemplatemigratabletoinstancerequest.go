package templatemigratableto

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTemplateMigratableToInstanceRequest struct {
	Description     nullable.Type[string]                   `json:"description,omitempty"`
	DisplayName     nullable.Type[string]                   `json:"displayName,omitempty"`
	RoleScopeTagIds *[]string                               `json:"roleScopeTagIds,omitempty"`
	SettingsDelta   *[]beta.DeviceManagementSettingInstance `json:"settingsDelta,omitempty"`
}
