package accessreviewdecisioninstancestagedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInstanceStageDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInstanceStageDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInstanceStageDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninstancestagedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInstanceStageDecisionClient: %+v", err)
	}

	return &AccessReviewDecisionInstanceStageDecisionClient{
		Client: client,
	}, nil
}
