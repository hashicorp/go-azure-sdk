package v2021_06_30

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/customoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fieresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fileresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/get"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/patch"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/projectresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/serviceresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/servicetaskresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/standardoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/taskresource"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CustomOperation     *customoperation.CustomOperationClient
	DELETE              *delete.DELETEClient
	FieResource         *fieresource.FieResourceClient
	FileResource        *fileresource.FileResourceClient
	GET                 *get.GETClient
	PATCH               *patch.PATCHClient
	POST                *post.POSTClient
	PUT                 *put.PUTClient
	ProjectResource     *projectresource.ProjectResourceClient
	ServiceResource     *serviceresource.ServiceResourceClient
	ServiceTaskResource *servicetaskresource.ServiceTaskResourceClient
	StandardOperation   *standardoperation.StandardOperationClient
	TaskResource        *taskresource.TaskResourceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	customOperationClient, err := customoperation.NewCustomOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomOperation client: %+v", err)
	}
	configureFunc(customOperationClient.Client)

	dELETEClient, err := delete.NewDELETEClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DELETE client: %+v", err)
	}
	configureFunc(dELETEClient.Client)

	fieResourceClient, err := fieresource.NewFieResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FieResource client: %+v", err)
	}
	configureFunc(fieResourceClient.Client)

	fileResourceClient, err := fileresource.NewFileResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileResource client: %+v", err)
	}
	configureFunc(fileResourceClient.Client)

	gETClient, err := get.NewGETClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GET client: %+v", err)
	}
	configureFunc(gETClient.Client)

	pATCHClient, err := patch.NewPATCHClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PATCH client: %+v", err)
	}
	configureFunc(pATCHClient.Client)

	pOSTClient, err := post.NewPOSTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building POST client: %+v", err)
	}
	configureFunc(pOSTClient.Client)

	pUTClient, err := put.NewPUTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PUT client: %+v", err)
	}
	configureFunc(pUTClient.Client)

	projectResourceClient, err := projectresource.NewProjectResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProjectResource client: %+v", err)
	}
	configureFunc(projectResourceClient.Client)

	serviceResourceClient, err := serviceresource.NewServiceResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceResource client: %+v", err)
	}
	configureFunc(serviceResourceClient.Client)

	serviceTaskResourceClient, err := servicetaskresource.NewServiceTaskResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceTaskResource client: %+v", err)
	}
	configureFunc(serviceTaskResourceClient.Client)

	standardOperationClient, err := standardoperation.NewStandardOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandardOperation client: %+v", err)
	}
	configureFunc(standardOperationClient.Client)

	taskResourceClient, err := taskresource.NewTaskResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TaskResource client: %+v", err)
	}
	configureFunc(taskResourceClient.Client)

	return &Client{
		CustomOperation:     customOperationClient,
		DELETE:              dELETEClient,
		FieResource:         fieResourceClient,
		FileResource:        fileResourceClient,
		GET:                 gETClient,
		PATCH:               pATCHClient,
		POST:                pOSTClient,
		PUT:                 pUTClient,
		ProjectResource:     projectResourceClient,
		ServiceResource:     serviceResourceClient,
		ServiceTaskResource: serviceTaskResourceClient,
		StandardOperation:   standardOperationClient,
		TaskResource:        taskResourceClient,
	}, nil
}
