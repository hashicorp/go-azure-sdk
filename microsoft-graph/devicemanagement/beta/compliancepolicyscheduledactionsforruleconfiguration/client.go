package compliancepolicyscheduledactionsforruleconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicyScheduledActionsForRuleConfigurationClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicyScheduledActionsForRuleConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicyScheduledActionsForRuleConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicyscheduledactionsforruleconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicyScheduledActionsForRuleConfigurationClient: %+v", err)
	}

	return &CompliancePolicyScheduledActionsForRuleConfigurationClient{
		Client: client,
	}, nil
}
