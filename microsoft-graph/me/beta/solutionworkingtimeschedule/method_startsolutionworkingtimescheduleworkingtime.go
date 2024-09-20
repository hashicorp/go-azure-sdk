package solutionworkingtimeschedule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartSolutionWorkingTimeScheduleWorkingTimeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultStartSolutionWorkingTimeScheduleWorkingTimeOperationOptions() StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions {
	return StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions{}
}

func (o StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StartSolutionWorkingTimeScheduleWorkingTime - Invoke action startWorkingTime. Triggers the policies associated with
// the end of working hours.
func (c SolutionWorkingTimeScheduleClient) StartSolutionWorkingTimeScheduleWorkingTime(ctx context.Context, options StartSolutionWorkingTimeScheduleWorkingTimeOperationOptions) (result StartSolutionWorkingTimeScheduleWorkingTimeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/solutions/workingTimeSchedule/startWorkingTime",
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

	return
}
