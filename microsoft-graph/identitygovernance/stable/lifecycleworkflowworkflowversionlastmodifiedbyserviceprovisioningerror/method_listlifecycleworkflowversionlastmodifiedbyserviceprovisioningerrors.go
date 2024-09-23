package lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror

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

type ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions() ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions {
	return ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions{}
}

func (o ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrors - Get serviceProvisioningErrors property value.
// Errors published by a federated service describing a nontransient, service-specific error regarding the properties or
// link from a user object. Supports $filter (eq, not, for isResolved and serviceInstance).
func (c LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrors(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, options ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions) (result ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/lastModifiedBy/serviceProvisioningErrors", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, options ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions) (ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient) ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, options ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrors(ctx, id, options)
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

	result = ListLifecycleWorkflowVersionLastModifiedByServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
