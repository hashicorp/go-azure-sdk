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

type CreateTeamUnarchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateTeamUnarchive - Invoke action unarchive. Restore an archived team. This API restores users' ability to send
// messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API.
// Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which might
// occur subsequent to a response from this API.
func (c TeamClient) CreateTeamUnarchive(ctx context.Context, id stable.GroupId) (result CreateTeamUnarchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/team/unarchive", id.ID()),
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
