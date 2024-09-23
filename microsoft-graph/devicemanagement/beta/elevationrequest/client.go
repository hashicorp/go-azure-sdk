package elevationrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElevationRequestClient struct {
	Client *msgraph.Client
}

func NewElevationRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*ElevationRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "elevationrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElevationRequestClient: %+v", err)
	}

	return &ElevationRequestClient{
		Client: client,
	}, nil
}
