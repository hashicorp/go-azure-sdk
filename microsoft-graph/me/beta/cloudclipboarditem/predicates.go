package cloudclipboarditem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CloudClipboardItemOperationPredicate struct {
}

func (p CloudClipboardItemOperationPredicate) Matches(input beta.CloudClipboardItem) bool {

	return true
}
