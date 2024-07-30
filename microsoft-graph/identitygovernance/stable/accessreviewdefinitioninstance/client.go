package accessreviewdefinitioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceClient{
		Client: client,
	}, nil
}
