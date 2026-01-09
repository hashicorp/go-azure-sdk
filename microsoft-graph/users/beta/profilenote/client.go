package profilenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileNoteClient struct {
	Client *msgraph.Client
}

func NewProfileNoteClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileNoteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileNoteClient: %+v", err)
	}

	return &ProfileNoteClient{
		Client: client,
	}, nil
}
