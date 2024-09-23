package mobilethreatdefenseconnector

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type MobileThreatDefenseConnectorOperationPredicate struct {
}

func (p MobileThreatDefenseConnectorOperationPredicate) Matches(input beta.MobileThreatDefenseConnector) bool {

	return true
}
