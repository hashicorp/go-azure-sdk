package calendargroupcalendarcalendarpermission

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type CalendarPermissionOperationPredicate struct {
}

func (p CalendarPermissionOperationPredicate) Matches(input stable.CalendarPermission) bool {

	return true
}
