package configurationcategory

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementConfigurationCategoryOperationPredicate struct {
}

func (p DeviceManagementConfigurationCategoryOperationPredicate) Matches(input beta.DeviceManagementConfigurationCategory) bool {

	return true
}
