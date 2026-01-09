package accessreviewdefinitioninstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageClient{
		Client: client,
	}, nil
}
