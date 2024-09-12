package lifecycleworkflowdeleteditemworkflowversionlastmodifiedbyserviceprovisioningerror

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

type GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCount - Get the number of the
// resource
func (c LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient) GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCount(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionId, options GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastModifiedBy/serviceProvisioningErrors/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
