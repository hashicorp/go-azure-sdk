package v2021_06_30

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	customOperationClient := customoperation.NewCustomOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&customOperationClient.Client)

	dELETEClient := delete.NewDELETEClientWithBaseURI(endpoint)
	configureAuthFunc(&dELETEClient.Client)

	fieResourceClient := fieresource.NewFieResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&fieResourceClient.Client)

	fileResourceClient := fileresource.NewFileResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&fileResourceClient.Client)

	gETClient := get.NewGETClientWithBaseURI(endpoint)
	configureAuthFunc(&gETClient.Client)

	pATCHClient := patch.NewPATCHClientWithBaseURI(endpoint)
	configureAuthFunc(&pATCHClient.Client)

	pOSTClient := post.NewPOSTClientWithBaseURI(endpoint)
	configureAuthFunc(&pOSTClient.Client)

	pUTClient := put.NewPUTClientWithBaseURI(endpoint)
	configureAuthFunc(&pUTClient.Client)

	projectResourceClient := projectresource.NewProjectResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&projectResourceClient.Client)

	serviceResourceClient := serviceresource.NewServiceResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceResourceClient.Client)

	serviceTaskResourceClient := servicetaskresource.NewServiceTaskResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceTaskResourceClient.Client)

	standardOperationClient := standardoperation.NewStandardOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&standardOperationClient.Client)

	taskResourceClient := taskresource.NewTaskResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&taskResourceClient.Client)

	return Client{
		CustomOperation:     &customOperationClient,
		DELETE:              &dELETEClient,
		FieResource:         &fieResourceClient,
		FileResource:        &fileResourceClient,
		GET:                 &gETClient,
		PATCH:               &pATCHClient,
		POST:                &pOSTClient,
		PUT:                 &pUTClient,
		ProjectResource:     &projectResourceClient,
		ServiceResource:     &serviceResourceClient,
		ServiceTaskResource: &serviceTaskResourceClient,
		StandardOperation:   &standardOperationClient,
		TaskResource:        &taskResourceClient,
	}
}
