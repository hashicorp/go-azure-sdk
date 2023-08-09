package attestations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsClient struct {
	Client *resourcemanager.Client
}

func NewAttestationsClientWithBaseURI(sdkApi sdkEnv.Api) (*AttestationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "attestations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttestationsClient: %+v", err)
	}

	return &AttestationsClient{
		Client: client,
	}, nil
}
