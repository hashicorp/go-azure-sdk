package functionenvelopeoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionEnvelopeOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewFunctionEnvelopeOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*FunctionEnvelopeOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "functionenvelopeoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FunctionEnvelopeOperationGroupClient: %+v", err)
	}

	return &FunctionEnvelopeOperationGroupClient{
		Client: client,
	}, nil
}
