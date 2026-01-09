package intunebrandingprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandingProfileClient struct {
	Client *msgraph.Client
}

func NewIntuneBrandingProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*IntuneBrandingProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intunebrandingprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntuneBrandingProfileClient: %+v", err)
	}

	return &IntuneBrandingProfileClient{
		Client: client,
	}, nil
}
