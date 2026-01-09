package accessreviewdefinitioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionClient{
		Client: client,
	}, nil
}
