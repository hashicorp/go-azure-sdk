package onlinemeetingtranscript

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type CallTranscriptOperationPredicate struct {
}

func (p CallTranscriptOperationPredicate) Matches(input stable.CallTranscript) bool {

	return true
}
