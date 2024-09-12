package configurationpolicytemplate

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementConfigurationPolicyTemplateOperationPredicate struct {
}

func (p DeviceManagementConfigurationPolicyTemplateOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicyTemplate) bool {

	return true
}
