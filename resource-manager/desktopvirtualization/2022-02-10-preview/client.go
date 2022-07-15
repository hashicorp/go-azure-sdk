package v2022_02_10_preview

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/desktop"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/msiximage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/msixpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/privatelink"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/scalingplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/sessionhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/startmenuitem"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/usersession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/workspace"
)

type Client struct {
	Application      *application.ApplicationClient
	ApplicationGroup *applicationgroup.ApplicationGroupClient
	Desktop          *desktop.DesktopClient
	HostPool         *hostpool.HostPoolClient
	MSIXPackage      *msixpackage.MSIXPackageClient
	MsixImage        *msiximage.MsixImageClient
	PrivateLink      *privatelink.PrivateLinkClient
	ScalingPlan      *scalingplan.ScalingPlanClient
	SessionHost      *sessionhost.SessionHostClient
	StartMenuItem    *startmenuitem.StartMenuItemClient
	UserSession      *usersession.UserSessionClient
	Workspace        *workspace.WorkspaceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationClient := application.NewApplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationClient.Client)

	applicationGroupClient := applicationgroup.NewApplicationGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationGroupClient.Client)

	desktopClient := desktop.NewDesktopClientWithBaseURI(endpoint)
	configureAuthFunc(&desktopClient.Client)

	hostPoolClient := hostpool.NewHostPoolClientWithBaseURI(endpoint)
	configureAuthFunc(&hostPoolClient.Client)

	mSIXPackageClient := msixpackage.NewMSIXPackageClientWithBaseURI(endpoint)
	configureAuthFunc(&mSIXPackageClient.Client)

	msixImageClient := msiximage.NewMsixImageClientWithBaseURI(endpoint)
	configureAuthFunc(&msixImageClient.Client)

	privateLinkClient := privatelink.NewPrivateLinkClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkClient.Client)

	scalingPlanClient := scalingplan.NewScalingPlanClientWithBaseURI(endpoint)
	configureAuthFunc(&scalingPlanClient.Client)

	sessionHostClient := sessionhost.NewSessionHostClientWithBaseURI(endpoint)
	configureAuthFunc(&sessionHostClient.Client)

	startMenuItemClient := startmenuitem.NewStartMenuItemClientWithBaseURI(endpoint)
	configureAuthFunc(&startMenuItemClient.Client)

	userSessionClient := usersession.NewUserSessionClientWithBaseURI(endpoint)
	configureAuthFunc(&userSessionClient.Client)

	workspaceClient := workspace.NewWorkspaceClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceClient.Client)

	return Client{
		Application:      &applicationClient,
		ApplicationGroup: &applicationGroupClient,
		Desktop:          &desktopClient,
		HostPool:         &hostPoolClient,
		MSIXPackage:      &mSIXPackageClient,
		MsixImage:        &msixImageClient,
		PrivateLink:      &privateLinkClient,
		ScalingPlan:      &scalingPlanClient,
		SessionHost:      &sessionHostClient,
		StartMenuItem:    &startMenuItemClient,
		UserSession:      &userSessionClient,
		Workspace:        &workspaceClient,
	}
}
