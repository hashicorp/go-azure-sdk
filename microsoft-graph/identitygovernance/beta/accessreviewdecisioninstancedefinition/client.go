package accessreviewdecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInstanceDefinitionClient: %+v", err)
	}

	return &AccessReviewDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
