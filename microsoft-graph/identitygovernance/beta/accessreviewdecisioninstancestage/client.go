package accessreviewdecisioninstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInstanceStageClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInstanceStageClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInstanceStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInstanceStageClient: %+v", err)
	}

	return &AccessReviewDecisionInstanceStageClient{
		Client: client,
	}, nil
}
