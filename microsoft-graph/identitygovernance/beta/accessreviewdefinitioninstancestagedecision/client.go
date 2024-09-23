package accessreviewdefinitioninstancestagedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestagedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageDecisionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageDecisionClient{
		Client: client,
	}, nil
}
