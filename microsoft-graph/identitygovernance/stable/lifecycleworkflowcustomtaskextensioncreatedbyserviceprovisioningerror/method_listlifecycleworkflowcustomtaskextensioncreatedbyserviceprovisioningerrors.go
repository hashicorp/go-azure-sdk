package lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror

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

type ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrors ...
func (c LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient) ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrors(ctx context.Context, id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) (result ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/createdBy/serviceProvisioningErrors", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.ServiceProvisioningError `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient) ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) (ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient) ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrors(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
