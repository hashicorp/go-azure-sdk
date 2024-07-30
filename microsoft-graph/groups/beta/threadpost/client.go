package threadpost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostClient struct {
	Client *msgraph.Client
}

func NewThreadPostClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostClient: %+v", err)
	}

	return &ThreadPostClient{
		Client: client,
	}, nil
}
