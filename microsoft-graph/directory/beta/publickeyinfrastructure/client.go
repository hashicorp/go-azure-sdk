package publickeyinfrastructure

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicKeyInfrastructureClient struct {
	Client *msgraph.Client
}

func NewPublicKeyInfrastructureClientWithBaseURI(sdkApi sdkEnv.Api) (*PublicKeyInfrastructureClient, error) {
	client, err := msgraph.NewClient(sdkApi, "publickeyinfrastructure", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublicKeyInfrastructureClient: %+v", err)
	}

	return &PublicKeyInfrastructureClient{
		Client: client,
	}, nil
}
