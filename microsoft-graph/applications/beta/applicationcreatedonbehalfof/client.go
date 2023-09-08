package applicationcreatedonbehalfof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationCreatedOnBehalfOfClient struct {
	Client *msgraph.Client
}

func NewApplicationCreatedOnBehalfOfClientWithBaseURI(api sdkEnv.Api) (*ApplicationCreatedOnBehalfOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationcreatedonbehalfof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationCreatedOnBehalfOfClient: %+v", err)
	}

	return &ApplicationCreatedOnBehalfOfClient{
		Client: client,
	}, nil
}
