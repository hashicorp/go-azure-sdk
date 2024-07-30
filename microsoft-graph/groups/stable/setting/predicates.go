package setting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type GroupSettingOperationPredicate struct {
}

func (p GroupSettingOperationPredicate) Matches(input stable.GroupSetting) bool {

	return true
}
