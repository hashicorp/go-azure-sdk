package teampermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewTeamPermissionGrantClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPermissionGrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teampermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPermissionGrantClient: %+v", err)
	}

	return &TeamPermissionGrantClient{
		Client: client,
	}, nil
}
