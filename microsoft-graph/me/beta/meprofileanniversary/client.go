package meprofileanniversary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileAnniversaryClient struct {
	Client *msgraph.Client
}

func NewMeProfileAnniversaryClientWithBaseURI(api sdkEnv.Api) (*MeProfileAnniversaryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileanniversary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileAnniversaryClient: %+v", err)
	}

	return &MeProfileAnniversaryClient{
		Client: client,
	}, nil
}
