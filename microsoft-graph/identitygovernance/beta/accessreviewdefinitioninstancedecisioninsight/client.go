package accessreviewdefinitioninstancedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInsightClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInsightClient{
		Client: client,
	}, nil
}
