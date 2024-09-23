package accessreviewdecisioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInstanceDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInstanceDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInstanceDecisionClient: %+v", err)
	}

	return &AccessReviewDecisionInstanceDecisionClient{
		Client: client,
	}, nil
}
