package profileskill

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileSkillClient struct {
	Client *msgraph.Client
}

func NewProfileSkillClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileSkillClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileskill", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileSkillClient: %+v", err)
	}

	return &ProfileSkillClient{
		Client: client,
	}, nil
}
