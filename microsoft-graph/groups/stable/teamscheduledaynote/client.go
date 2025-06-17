package teamscheduledaynote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleDayNoteClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleDayNoteClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleDayNoteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduledaynote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleDayNoteClient: %+v", err)
	}

	return &TeamScheduleDayNoteClient{
		Client: client,
	}, nil
}
