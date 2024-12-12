package v2024_04_08_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/activesessionhostconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/appattachpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/appattachpackageinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/desktop"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/msiximage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/msixpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/privatelink"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpersonalschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpooledschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostmanagement"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostmanagements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/startmenuitem"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/usersession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/workspace"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActiveSessionHostConfiguration *activesessionhostconfiguration.ActiveSessionHostConfigurationClient
	AppAttachPackage               *appattachpackage.AppAttachPackageClient
	AppAttachPackageInfo           *appattachpackageinfo.AppAttachPackageInfoClient
	Application                    *application.ApplicationClient
	ApplicationGroup               *applicationgroup.ApplicationGroupClient
	Desktop                        *desktop.DesktopClient
	HostPool                       *hostpool.HostPoolClient
	MSIXPackage                    *msixpackage.MSIXPackageClient
	MsixImage                      *msiximage.MsixImageClient
	PrivateLink                    *privatelink.PrivateLinkClient
	ScalingPlan                    *scalingplan.ScalingPlanClient
	ScalingPlanPersonalSchedule    *scalingplanpersonalschedule.ScalingPlanPersonalScheduleClient
	ScalingPlanPooledSchedule      *scalingplanpooledschedule.ScalingPlanPooledScheduleClient
	SessionHost                    *sessionhost.SessionHostClient
	SessionHostConfiguration       *sessionhostconfiguration.SessionHostConfigurationClient
	SessionHostManagement          *sessionhostmanagement.SessionHostManagementClient
	SessionHostManagements         *sessionhostmanagements.SessionHostManagementsClient
	StartMenuItem                  *startmenuitem.StartMenuItemClient
	UserSession                    *usersession.UserSessionClient
	Workspace                      *workspace.WorkspaceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activeSessionHostConfigurationClient, err := activesessionhostconfiguration.NewActiveSessionHostConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActiveSessionHostConfiguration client: %+v", err)
	}
	configureFunc(activeSessionHostConfigurationClient.Client)

	appAttachPackageClient, err := appattachpackage.NewAppAttachPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppAttachPackage client: %+v", err)
	}
	configureFunc(appAttachPackageClient.Client)

	appAttachPackageInfoClient, err := appattachpackageinfo.NewAppAttachPackageInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppAttachPackageInfo client: %+v", err)
	}
	configureFunc(appAttachPackageInfoClient.Client)

	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationGroupClient, err := applicationgroup.NewApplicationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGroup client: %+v", err)
	}
	configureFunc(applicationGroupClient.Client)

	desktopClient, err := desktop.NewDesktopClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Desktop client: %+v", err)
	}
	configureFunc(desktopClient.Client)

	hostPoolClient, err := hostpool.NewHostPoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HostPool client: %+v", err)
	}
	configureFunc(hostPoolClient.Client)

	mSIXPackageClient, err := msixpackage.NewMSIXPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSIXPackage client: %+v", err)
	}
	configureFunc(mSIXPackageClient.Client)

	msixImageClient, err := msiximage.NewMsixImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MsixImage client: %+v", err)
	}
	configureFunc(msixImageClient.Client)

	privateLinkClient, err := privatelink.NewPrivateLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLink client: %+v", err)
	}
	configureFunc(privateLinkClient.Client)

	scalingPlanClient, err := scalingplan.NewScalingPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlan client: %+v", err)
	}
	configureFunc(scalingPlanClient.Client)

	scalingPlanPersonalScheduleClient, err := scalingplanpersonalschedule.NewScalingPlanPersonalScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlanPersonalSchedule client: %+v", err)
	}
	configureFunc(scalingPlanPersonalScheduleClient.Client)

	scalingPlanPooledScheduleClient, err := scalingplanpooledschedule.NewScalingPlanPooledScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlanPooledSchedule client: %+v", err)
	}
	configureFunc(scalingPlanPooledScheduleClient.Client)

	sessionHostClient, err := sessionhost.NewSessionHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHost client: %+v", err)
	}
	configureFunc(sessionHostClient.Client)

	sessionHostConfigurationClient, err := sessionhostconfiguration.NewSessionHostConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHostConfiguration client: %+v", err)
	}
	configureFunc(sessionHostConfigurationClient.Client)

	sessionHostManagementClient, err := sessionhostmanagement.NewSessionHostManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHostManagement client: %+v", err)
	}
	configureFunc(sessionHostManagementClient.Client)

	sessionHostManagementsClient, err := sessionhostmanagements.NewSessionHostManagementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHostManagements client: %+v", err)
	}
	configureFunc(sessionHostManagementsClient.Client)

	startMenuItemClient, err := startmenuitem.NewStartMenuItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StartMenuItem client: %+v", err)
	}
	configureFunc(startMenuItemClient.Client)

	userSessionClient, err := usersession.NewUserSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSession client: %+v", err)
	}
	configureFunc(userSessionClient.Client)

	workspaceClient, err := workspace.NewWorkspaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspace client: %+v", err)
	}
	configureFunc(workspaceClient.Client)

	return &Client{
		ActiveSessionHostConfiguration: activeSessionHostConfigurationClient,
		AppAttachPackage:               appAttachPackageClient,
		AppAttachPackageInfo:           appAttachPackageInfoClient,
		Application:                    applicationClient,
		ApplicationGroup:               applicationGroupClient,
		Desktop:                        desktopClient,
		HostPool:                       hostPoolClient,
		MSIXPackage:                    mSIXPackageClient,
		MsixImage:                      msixImageClient,
		PrivateLink:                    privateLinkClient,
		ScalingPlan:                    scalingPlanClient,
		ScalingPlanPersonalSchedule:    scalingPlanPersonalScheduleClient,
		ScalingPlanPooledSchedule:      scalingPlanPooledScheduleClient,
		SessionHost:                    sessionHostClient,
		SessionHostConfiguration:       sessionHostConfigurationClient,
		SessionHostManagement:          sessionHostManagementClient,
		SessionHostManagements:         sessionHostManagementsClient,
		StartMenuItem:                  startMenuItemClient,
		UserSession:                    userSessionClient,
		Workspace:                      workspaceClient,
	}, nil
}
