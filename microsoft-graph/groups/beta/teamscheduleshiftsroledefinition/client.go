package teamscheduleshiftsroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleShiftsRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleShiftsRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleShiftsRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduleshiftsroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleShiftsRoleDefinitionClient: %+v", err)
	}

	return &TeamScheduleShiftsRoleDefinitionClient{
		Client: client,
	}, nil
}
