package inferenceclassificationoverride

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationOverrideClient struct {
	Client *msgraph.Client
}

func NewInferenceClassificationOverrideClientWithBaseURI(sdkApi sdkEnv.Api) (*InferenceClassificationOverrideClient, error) {
	client, err := msgraph.NewClient(sdkApi, "inferenceclassificationoverride", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InferenceClassificationOverrideClient: %+v", err)
	}

	return &InferenceClassificationOverrideClient{
		Client: client,
	}, nil
}
