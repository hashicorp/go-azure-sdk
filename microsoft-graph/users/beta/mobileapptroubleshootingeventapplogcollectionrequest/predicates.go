package mobileapptroubleshootingeventapplogcollectionrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AppLogCollectionRequestOperationPredicate struct {
}

func (p AppLogCollectionRequestOperationPredicate) Matches(input beta.AppLogCollectionRequest) bool {

	return true
}
