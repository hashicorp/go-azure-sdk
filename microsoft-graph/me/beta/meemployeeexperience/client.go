package meemployeeexperience

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEmployeeExperienceClient struct {
	Client *msgraph.Client
}

func NewMeEmployeeExperienceClientWithBaseURI(api sdkEnv.Api) (*MeEmployeeExperienceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meemployeeexperience", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEmployeeExperienceClient: %+v", err)
	}

	return &MeEmployeeExperienceClient{
		Client: client,
	}, nil
}
