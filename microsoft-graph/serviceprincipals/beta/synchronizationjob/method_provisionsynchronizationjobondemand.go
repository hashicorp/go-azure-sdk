package synchronizationjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionSynchronizationJobOnDemandOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.StringKeyStringValuePair
}

// ProvisionSynchronizationJobOnDemand - Invoke action provisionOnDemand. Select a user and provision the account
// on-demand. The rate limit for this API is 5 requests per 10 seconds. No user or group will be provisioned on-demand
// that would not have been provisioned through the regular provisioning cycles.
func (c SynchronizationJobClient) ProvisionSynchronizationJobOnDemand(ctx context.Context, id beta.ServicePrincipalIdSynchronizationJobId, input ProvisionSynchronizationJobOnDemandRequest) (result ProvisionSynchronizationJobOnDemandOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/provisionOnDemand", id.ID()),
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

	var model beta.StringKeyStringValuePair
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
