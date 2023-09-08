package meemployeeexperiencelearningcourseactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEmployeeExperienceLearningCourseActivityClient struct {
	Client *msgraph.Client
}

func NewMeEmployeeExperienceLearningCourseActivityClientWithBaseURI(api sdkEnv.Api) (*MeEmployeeExperienceLearningCourseActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meemployeeexperiencelearningcourseactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEmployeeExperienceLearningCourseActivityClient: %+v", err)
	}

	return &MeEmployeeExperienceLearningCourseActivityClient{
		Client: client,
	}, nil
}
