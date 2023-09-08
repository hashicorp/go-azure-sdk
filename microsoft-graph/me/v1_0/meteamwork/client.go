package meteamwork

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkClient: %+v", err)
	}

	return &MeTeamworkClient{
		Client: client,
	}, nil
}
