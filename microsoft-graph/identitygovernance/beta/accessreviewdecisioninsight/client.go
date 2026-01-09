package accessreviewdecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInsightClient: %+v", err)
	}

	return &AccessReviewDecisionInsightClient{
		Client: client,
	}, nil
}
