package teamtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateClient struct {
	Client *msgraph.Client
}

func NewTeamTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamTemplateClient: %+v", err)
	}

	return &TeamTemplateClient{
		Client: client,
	}, nil
}
