package v2021_10_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/image"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/lab"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/labplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/virtualmachine"
)

type Client struct {
	Image          *image.ImageClient
	Lab            *lab.LabClient
	LabPlan        *labplan.LabPlanClient
	Schedule       *schedule.ScheduleClient
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

	userClient := user.NewUserClientWithBaseURI(endpoint)
	configureAuthFunc(&userClient.Client)

	virtualMachineClient := virtualmachine.NewVirtualMachineClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineClient.Client)

	return Client{
		Image:          &imageClient,
		Lab:            &labClient,
		LabPlan:        &labPlanClient,
		Schedule:       &scheduleClient,
		User:           &userClient,
		VirtualMachine: &virtualMachineClient,
	}
}
