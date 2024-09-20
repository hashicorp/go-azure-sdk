package teamschedule

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

type ShareTeamScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ShareTeamScheduleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultShareTeamScheduleOperationOptions() ShareTeamScheduleOperationOptions {
	return ShareTeamScheduleOperationOptions{}
}

func (o ShareTeamScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ShareTeamScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ShareTeamScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ShareTeamSchedule - Invoke action share. Share a schedule time range with schedule members. This action makes the
// collections of shift, openshift and timeOff items in the specified time range of the schedule viewable by the
// specified team members, including employees and managers. Each shift, openshift and timeOff instance in a schedule
// supports a draft version and a shared version of the item. The draft version is viewable by only managers, and the
// shared version is viewable by employees and managers. For each shift, openshift and timeOff instance in the specified
// time range, the share action updates the shared version from the draft version, so that in addition to managers,
// employees can also view the most current information about the item. The notifyTeam parameter further specifies which
// employees can view the item.
func (c TeamScheduleClient) ShareTeamSchedule(ctx context.Context, id stable.GroupId, input ShareTeamScheduleRequest, options ShareTeamScheduleOperationOptions) (result ShareTeamScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/schedule/share", id.ID()),
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
