package homerealmdiscoverypolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type HomeRealmDiscoveryPolicyOperationPredicate struct {
}

func (p HomeRealmDiscoveryPolicyOperationPredicate) Matches(input stable.HomeRealmDiscoveryPolicy) bool {

	return true
}
