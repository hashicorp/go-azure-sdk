package group

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

type AssignLicenseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Group
}

// AssignLicense - Invoke action assignLicense. Add or remove licenses on the group. Licenses assigned to the group will
// be assigned to all users in the group. To learn more about group-based licensing, see What is group-based licensing
// in Microsoft Entra ID. To get the subscriptions available in the directory, perform a GET subscribedSkus request.
func (c GroupClient) AssignLicense(ctx context.Context, id stable.GroupId, input AssignLicenseRequest) (result AssignLicenseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/assignLicense", id.ID()),
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

	var model stable.Group
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
