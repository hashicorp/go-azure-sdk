package joinedteamincomingchannel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ChannelOperationPredicate struct {
}

func (p ChannelOperationPredicate) Matches(input stable.Channel) bool {

	return true
}
