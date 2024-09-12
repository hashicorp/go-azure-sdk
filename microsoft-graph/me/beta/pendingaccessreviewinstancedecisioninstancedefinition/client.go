package pendingaccessreviewinstancedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceDecisionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceDecisionInstanceDefinitionClient: %+v", err)
	}

	return &PendingAccessReviewInstanceDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
