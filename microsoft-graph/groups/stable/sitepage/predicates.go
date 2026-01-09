package sitepage

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type BaseSitePageOperationPredicate struct {
}

func (p BaseSitePageOperationPredicate) Matches(input stable.BaseSitePage) bool {

	return true
}

type SitePageOperationPredicate struct {
}

func (p SitePageOperationPredicate) Matches(input stable.SitePage) bool {

	return true
}
