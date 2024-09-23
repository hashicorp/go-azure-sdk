package accessreviewdefinitioninstancedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInstanceClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInstanceClient{
		Client: client,
	}, nil
}
