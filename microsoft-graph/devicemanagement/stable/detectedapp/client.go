package detectedapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedAppClient struct {
	Client *msgraph.Client
}

func NewDetectedAppClientWithBaseURI(sdkApi sdkEnv.Api) (*DetectedAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "detectedapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DetectedAppClient: %+v", err)
	}

	return &DetectedAppClient{
		Client: client,
	}, nil
}
