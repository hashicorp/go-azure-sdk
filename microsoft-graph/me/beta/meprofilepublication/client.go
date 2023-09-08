package meprofilepublication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfilePublicationClient struct {
	Client *msgraph.Client
}

func NewMeProfilePublicationClientWithBaseURI(api sdkEnv.Api) (*MeProfilePublicationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilepublication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfilePublicationClient: %+v", err)
	}

	return &MeProfilePublicationClient{
		Client: client,
	}, nil
}
