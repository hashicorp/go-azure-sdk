package recommendedactionsessions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorsStartRecommendedActionSessionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type AdvisorsStartRecommendedActionSessionOperationOptions struct {
	DatabaseName *string
}

func DefaultAdvisorsStartRecommendedActionSessionOperationOptions() AdvisorsStartRecommendedActionSessionOperationOptions {
	return AdvisorsStartRecommendedActionSessionOperationOptions{}
}

func (o AdvisorsStartRecommendedActionSessionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AdvisorsStartRecommendedActionSessionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AdvisorsStartRecommendedActionSessionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatabaseName != nil {
		out.Append("databaseName", fmt.Sprintf("%v", *o.DatabaseName))
	}
	return &out
}

// AdvisorsStartRecommendedActionSession ...
func (c RecommendedActionSessionsClient) AdvisorsStartRecommendedActionSession(ctx context.Context, id AdvisorId, options AdvisorsStartRecommendedActionSessionOperationOptions) (result AdvisorsStartRecommendedActionSessionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/recommendedActionSessions", id.ID()),
		OptionsObject: options,
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

// AdvisorsStartRecommendedActionSessionThenPoll performs AdvisorsStartRecommendedActionSession then polls until it's completed
func (c RecommendedActionSessionsClient) AdvisorsStartRecommendedActionSessionThenPoll(ctx context.Context, id AdvisorId, options AdvisorsStartRecommendedActionSessionOperationOptions) error {
	result, err := c.AdvisorsStartRecommendedActionSession(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing AdvisorsStartRecommendedActionSession: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AdvisorsStartRecommendedActionSession: %+v", err)
	}

	return nil
}
