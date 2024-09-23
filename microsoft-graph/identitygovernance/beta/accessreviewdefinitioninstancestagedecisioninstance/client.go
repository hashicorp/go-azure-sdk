package accessreviewdefinitioninstancestagedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageDecisionInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageDecisionInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestagedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageDecisionInstanceClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageDecisionInstanceClient{
		Client: client,
	}, nil
}
