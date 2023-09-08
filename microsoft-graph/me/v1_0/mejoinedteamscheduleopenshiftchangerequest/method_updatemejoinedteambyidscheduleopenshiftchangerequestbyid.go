package mejoinedteamscheduleopenshiftchangerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMeJoinedTeamByIdScheduleOpenShiftChangeRequestByIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *models.OpenShiftChangeRequest
}

// UpdateMeJoinedTeamByIdScheduleOpenShiftChangeRequestById ...
func (c MeJoinedTeamScheduleOpenShiftChangeRequestClient) UpdateMeJoinedTeamByIdScheduleOpenShiftChangeRequestById(ctx context.Context, id MeJoinedTeamScheduleOpenShiftChangeRequestId, input models.OpenShiftChangeRequest) (result UpdateMeJoinedTeamByIdScheduleOpenShiftChangeRequestByIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
