package partner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerClient struct {
	Client *msgraph.Client
}

func NewPartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerClient: %+v", err)
	}

	return &PartnerClient{
		Client: client,
	}, nil
}
