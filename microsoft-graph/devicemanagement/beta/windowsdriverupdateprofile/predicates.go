package windowsdriverupdateprofile

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type WindowsDriverUpdateProfileOperationPredicate struct {
}

func (p WindowsDriverUpdateProfileOperationPredicate) Matches(input beta.WindowsDriverUpdateProfile) bool {

	return true
}
