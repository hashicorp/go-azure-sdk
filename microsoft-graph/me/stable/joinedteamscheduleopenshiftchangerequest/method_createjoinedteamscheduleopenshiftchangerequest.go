package joinedteamscheduleopenshiftchangerequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedTeamScheduleOpenShiftChangeRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OpenShiftChangeRequest
}

type CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions() CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions {
	return CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions{}
}

func (o CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamScheduleOpenShiftChangeRequest - Create new navigation property to openShiftChangeRequests for me
func (c JoinedTeamScheduleOpenShiftChangeRequestClient) CreateJoinedTeamScheduleOpenShiftChangeRequest(ctx context.Context, id stable.MeJoinedTeamId, input stable.OpenShiftChangeRequest, options CreateJoinedTeamScheduleOpenShiftChangeRequestOperationOptions) (result CreateJoinedTeamScheduleOpenShiftChangeRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/schedule/openShiftChangeRequests", id.ID()),
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

	var model stable.OpenShiftChangeRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
