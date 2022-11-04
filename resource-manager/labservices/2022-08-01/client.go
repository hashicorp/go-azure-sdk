package v2022_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/image"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/lab"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/labplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/virtualmachine"
)

type Client struct {
	Image          *image.ImageClient
	Lab            *lab.LabClient
	LabPlan        *labplan.LabPlanClient
	Schedule       *schedule.ScheduleClient
	Skus           *skus.SkusClient
	Usages         *usages.UsagesClient
	User           *user.UserClient
	VirtualMachine *virtualmachine.VirtualMachineClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	imageClient := image.NewImageClientWithBaseURI(endpoint)
	configureAuthFunc(&imageClient.Client)

	labClient := lab.NewLabClientWithBaseURI(endpoint)
	configureAuthFunc(&labClient.Client)

	labPlanClient := labplan.NewLabPlanClientWithBaseURI(endpoint)
	configureAuthFunc(&labPlanClient.Client)

	scheduleClient := schedule.NewScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduleClient.Client)

	skusClient := skus.NewSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&skusClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	userClient := user.NewUserClientWithBaseURI(endpoint)
	configureAuthFunc(&userClient.Client)

	virtualMachineClient := virtualmachine.NewVirtualMachineClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineClient.Client)

	return Client{
		Image:          &imageClient,
		Lab:            &labClient,
		LabPlan:        &labPlanClient,
		Schedule:       &scheduleClient,
		Skus:           &skusClient,
		Usages:         &usagesClient,
		User:           &userClient,
		VirtualMachine: &virtualMachineClient,
	}
}
