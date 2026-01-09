package accessreviewdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionClient: %+v", err)
	}

	return &AccessReviewDefinitionClient{
		Client: client,
	}, nil
}
