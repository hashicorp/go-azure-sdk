package meprofilenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileNoteClient struct {
	Client *msgraph.Client
}

func NewMeProfileNoteClientWithBaseURI(api sdkEnv.Api) (*MeProfileNoteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileNoteClient: %+v", err)
	}

	return &MeProfileNoteClient{
		Client: client,
	}, nil
}
