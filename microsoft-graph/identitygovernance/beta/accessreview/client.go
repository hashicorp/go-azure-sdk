package accessreview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewClient struct {
	Client *msgraph.Client
}

func NewAccessReviewClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewClient: %+v", err)
	}

	return &AccessReviewClient{
		Client: client,
	}, nil
}
