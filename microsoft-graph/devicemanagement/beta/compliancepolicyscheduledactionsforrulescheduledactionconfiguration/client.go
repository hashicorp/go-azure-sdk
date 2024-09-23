package compliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicyscheduledactionsforrulescheduledactionconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient: %+v", err)
	}

	return &CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient{
		Client: client,
	}, nil
}
