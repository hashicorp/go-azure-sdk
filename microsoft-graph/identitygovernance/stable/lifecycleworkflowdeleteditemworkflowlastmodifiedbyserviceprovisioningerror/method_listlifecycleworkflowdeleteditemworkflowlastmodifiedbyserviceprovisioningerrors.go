package lifecycleworkflowdeleteditemworkflowlastmodifiedbyserviceprovisioningerror

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions() ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions {
	return ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions{}
}

func (o ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrors - Get serviceProvisioningErrors
// property value. Errors published by a federated service describing a nontransient, service-specific error regarding
// the properties or link from a user object. Supports $filter (eq, not, for isResolved and serviceInstance).
func (c LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrors(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions) (result ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/lastModifiedBy/serviceProvisioningErrors", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.ServiceProvisioningError, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalServiceProvisioningErrorImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ServiceProvisioningError (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions) (ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrors(ctx, id, options)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
