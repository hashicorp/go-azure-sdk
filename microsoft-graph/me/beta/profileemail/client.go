package profileemail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileEmailClient struct {
	Client *msgraph.Client
}

func NewProfileEmailClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileEmailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileemail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileEmailClient: %+v", err)
	}

	return &ProfileEmailClient{
		Client: client,
	}, nil
}
