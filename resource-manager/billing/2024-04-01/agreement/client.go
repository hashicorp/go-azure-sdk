package agreement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementClient struct {
	Client *resourcemanager.Client
}

func NewAgreementClientWithBaseURI(sdkApi sdkEnv.Api) (*AgreementClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "agreement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgreementClient: %+v", err)
	}

	return &AgreementClient{
		Client: client,
	}, nil
}
