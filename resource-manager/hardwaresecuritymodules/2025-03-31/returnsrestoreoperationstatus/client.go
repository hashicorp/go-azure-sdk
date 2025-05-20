package returnsrestoreoperationstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReturnsRestoreOperationStatusClient struct {
	Client *resourcemanager.Client
}

func NewReturnsRestoreOperationStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*ReturnsRestoreOperationStatusClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "returnsrestoreoperationstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReturnsRestoreOperationStatusClient: %+v", err)
	}

	return &ReturnsRestoreOperationStatusClient{
		Client: client,
	}, nil
}
