package createdonbehalfof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatedOnBehalfOfClient struct {
	Client *msgraph.Client
}

func NewCreatedOnBehalfOfClientWithBaseURI(sdkApi sdkEnv.Api) (*CreatedOnBehalfOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "createdonbehalfof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CreatedOnBehalfOfClient: %+v", err)
	}

	return &CreatedOnBehalfOfClient{
		Client: client,
	}, nil
}
