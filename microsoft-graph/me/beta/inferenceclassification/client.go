package inferenceclassification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationClient struct {
	Client *msgraph.Client
}

func NewInferenceClassificationClientWithBaseURI(sdkApi sdkEnv.Api) (*InferenceClassificationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "inferenceclassification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InferenceClassificationClient: %+v", err)
	}

	return &InferenceClassificationClient{
		Client: client,
	}, nil
}
