package todolisttaskattachmentsession

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AttachmentSessionOperationPredicate struct {
}

func (p AttachmentSessionOperationPredicate) Matches(input stable.AttachmentSession) bool {

	return true
}
