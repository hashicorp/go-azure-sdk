package compliancepolicyscheduledactionsforrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicyScheduledActionsForRuleClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicyScheduledActionsForRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicyScheduledActionsForRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicyscheduledactionsforrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicyScheduledActionsForRuleClient: %+v", err)
	}

	return &CompliancePolicyScheduledActionsForRuleClient{
		Client: client,
	}, nil
}
