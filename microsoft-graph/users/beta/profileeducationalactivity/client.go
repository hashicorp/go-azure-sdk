package profileeducationalactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileEducationalActivityClient struct {
	Client *msgraph.Client
}

func NewProfileEducationalActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileEducationalActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileeducationalactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileEducationalActivityClient: %+v", err)
	}

	return &ProfileEducationalActivityClient{
		Client: client,
	}, nil
}
