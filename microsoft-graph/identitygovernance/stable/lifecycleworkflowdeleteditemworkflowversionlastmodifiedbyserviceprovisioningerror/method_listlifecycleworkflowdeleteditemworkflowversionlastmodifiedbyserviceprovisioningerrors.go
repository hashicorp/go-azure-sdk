package lifecycleworkflowdeleteditemworkflowversionlastmodifiedbyserviceprovisioningerror

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

type ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrors ...
func (c LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrors(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionId) (result ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/lastModifiedBy/serviceProvisioningErrors", id.ID()),
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

// ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionId) (ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionId, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrors(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
