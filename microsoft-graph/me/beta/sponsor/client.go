package sponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SponsorClient struct {
	Client *msgraph.Client
}

func NewSponsorClientWithBaseURI(sdkApi sdkEnv.Api) (*SponsorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SponsorClient: %+v", err)
	}

	return &SponsorClient{
		Client: client,
	}, nil
}
