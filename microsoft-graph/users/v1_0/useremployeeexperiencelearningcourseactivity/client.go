package useremployeeexperiencelearningcourseactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEmployeeExperienceLearningCourseActivityClient struct {
	Client *msgraph.Client
}

func NewUserEmployeeExperienceLearningCourseActivityClientWithBaseURI(api sdkEnv.Api) (*UserEmployeeExperienceLearningCourseActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useremployeeexperiencelearningcourseactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEmployeeExperienceLearningCourseActivityClient: %+v", err)
	}

	return &UserEmployeeExperienceLearningCourseActivityClient{
		Client: client,
	}, nil
}
