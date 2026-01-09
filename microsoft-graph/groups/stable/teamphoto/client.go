package teamphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPhotoClient struct {
	Client *msgraph.Client
}

func NewTeamPhotoClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPhotoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPhotoClient: %+v", err)
	}

	return &TeamPhotoClient{
		Client: client,
	}, nil
}
