package appserviceenvironmentresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsResumeOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Site
}

type AppServiceEnvironmentsResumeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Site
}

type AppServiceEnvironmentsResumeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AppServiceEnvironmentsResumeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AppServiceEnvironmentsResume ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsResume(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsResumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AppServiceEnvironmentsResumeCustomPager{},
		Path:       fmt.Sprintf("%s/resume", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// AppServiceEnvironmentsResumeThenPoll performs AppServiceEnvironmentsResume then polls until it's completed
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsResumeThenPoll(ctx context.Context, id commonids.AppServiceEnvironmentId) error {
	result, err := c.AppServiceEnvironmentsResume(ctx, id)
	if err != nil {
		return fmt.Errorf("performing AppServiceEnvironmentsResume: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AppServiceEnvironmentsResume: %+v", err)
	}

	return nil
}
