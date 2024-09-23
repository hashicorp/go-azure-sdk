package b2xuserflow

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowClient: %+v", err)
	}

	return &B2xUserFlowClient{
		Client: client,
	}, nil
}
