package v2018_04_19

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/customoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/get"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/patch"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/projectresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/serviceresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/standardoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/taskresource"
)

type Client struct {
	CustomOperation   *customoperation.CustomOperationClient
	DELETE            *delete.DELETEClient
	GET               *get.GETClient
	PATCH             *patch.PATCHClient
	POST              *post.POSTClient
	PUT               *put.PUTClient
	ProjectResource   *projectresource.ProjectResourceClient
	ServiceResource   *serviceresource.ServiceResourceClient
	StandardOperation *standardoperation.StandardOperationClient
	TaskResource      *taskresource.TaskResourceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	customOperationClient := customoperation.NewCustomOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&customOperationClient.Client)

	dELETEClient := delete.NewDELETEClientWithBaseURI(endpoint)
	configureAuthFunc(&dELETEClient.Client)

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

	standardOperationClient := standardoperation.NewStandardOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&standardOperationClient.Client)

	taskResourceClient := taskresource.NewTaskResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&taskResourceClient.Client)

	return Client{
		CustomOperation:   &customOperationClient,
		DELETE:            &dELETEClient,
		GET:               &gETClient,
		PATCH:             &pATCHClient,
		POST:              &pOSTClient,
		PUT:               &pUTClient,
		ProjectResource:   &projectResourceClient,
		ServiceResource:   &serviceResourceClient,
		StandardOperation: &standardOperationClient,
		TaskResource:      &taskResourceClient,
	}
}
