package devicedevicetemplate

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceTemplateOperationPredicate struct {
}

func (p DeviceTemplateOperationPredicate) Matches(input beta.DeviceTemplate) bool {

	return true
}
