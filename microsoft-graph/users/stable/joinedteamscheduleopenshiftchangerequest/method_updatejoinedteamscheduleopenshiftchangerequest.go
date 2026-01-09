package joinedteamscheduleopenshiftchangerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions() UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions {
	return UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions{}
}

func (o UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamScheduleOpenShiftChangeRequest - Update the navigation property openShiftChangeRequests in users
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) UpdateJoinedTeamScheduleOpenShiftChangeRequest(ctx context.Context, id stable.UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId, input stable.OpenShiftChangeRequest, options UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) (result UpdateJoinedTeamScheduleOpenShiftChangeRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
