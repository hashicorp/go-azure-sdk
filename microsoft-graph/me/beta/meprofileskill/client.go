package meprofileskill

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileSkillClient struct {
	Client *msgraph.Client
}

func NewMeProfileSkillClientWithBaseURI(api sdkEnv.Api) (*MeProfileSkillClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileskill", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileSkillClient: %+v", err)
	}

	return &MeProfileSkillClient{
		Client: client,
	}, nil
}
