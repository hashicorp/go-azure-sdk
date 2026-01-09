package solution

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionClient struct {
	Client *msgraph.Client
}

func NewSolutionClientWithBaseURI(sdkApi sdkEnv.Api) (*SolutionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "solution", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SolutionClient: %+v", err)
	}

	return &SolutionClient{
		Client: client,
	}, nil
}
