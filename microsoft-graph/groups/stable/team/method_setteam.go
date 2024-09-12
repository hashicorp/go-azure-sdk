package team

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

type SetTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// SetTeam - Create team from group. Create a new team under a group. In order to create a team, the group must have a
// least one owner. If the creation of the team call is delayed, you can retry the call up to three times before you
// have to wait for 15 minutes due to a propagation delay. If the group was created less than 15 minutes ago, the call
// might fail with a 404 error code due to replication delays. If the group was created less than 15 minutes ago, it's
// possible for a call to create a team to fail with a 404 error code, due to ongoing replication delays. The
// recommended pattern is to retry the Create team call three times, with a 10 second delay between calls.
func (c TeamClient) SetTeam(ctx context.Context, id stable.GroupId, input stable.Team) (result SetTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/team", id.ID()),
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
