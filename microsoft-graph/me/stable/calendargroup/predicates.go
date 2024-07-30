package calendargroup

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type CalendarGroupOperationPredicate struct {
}

func (p CalendarGroupOperationPredicate) Matches(input stable.CalendarGroup) bool {

	return true
}
