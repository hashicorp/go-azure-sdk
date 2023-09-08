package meinferenceclassification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInferenceClassificationClient struct {
	Client *msgraph.Client
}

func NewMeInferenceClassificationClientWithBaseURI(api sdkEnv.Api) (*MeInferenceClassificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinferenceclassification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInferenceClassificationClient: %+v", err)
	}

	return &MeInferenceClassificationClient{
		Client: client,
	}, nil
}
