package accessreviewdecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionClient: %+v", err)
	}

	return &AccessReviewDecisionClient{
		Client: client,
	}, nil
}
