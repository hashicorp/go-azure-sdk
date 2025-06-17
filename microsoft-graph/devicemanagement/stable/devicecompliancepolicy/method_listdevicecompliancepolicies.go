package devicecompliancepolicy

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

type ListDeviceCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceCompliancePolicy
}

type ListDeviceCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceCompliancePolicy
}

type ListDeviceCompliancePoliciesOperationOptions struct {
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

func DefaultListDeviceCompliancePoliciesOperationOptions() ListDeviceCompliancePoliciesOperationOptions {
	return ListDeviceCompliancePoliciesOperationOptions{}
}

func (o ListDeviceCompliancePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicies - List deviceCompliancePolicies. List properties and relationships of the
// deviceCompliancePolicy objects.
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePolicies(ctx context.Context, options ListDeviceCompliancePoliciesOperationOptions) (result ListDeviceCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePoliciesCustomPager{},
		Path:          "/deviceManagement/deviceCompliancePolicies",
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

	temp := make([]stable.DeviceCompliancePolicy, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDeviceCompliancePolicyImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DeviceCompliancePolicy (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDeviceCompliancePoliciesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePoliciesComplete(ctx context.Context, options ListDeviceCompliancePoliciesOperationOptions) (ListDeviceCompliancePoliciesCompleteResult, error) {
	return c.ListDeviceCompliancePoliciesCompleteMatchingPredicate(ctx, options, DeviceCompliancePolicyOperationPredicate{})
}

// ListDeviceCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, options ListDeviceCompliancePoliciesOperationOptions, predicate DeviceCompliancePolicyOperationPredicate) (result ListDeviceCompliancePoliciesCompleteResult, err error) {
	items := make([]stable.DeviceCompliancePolicy, 0)

	resp, err := c.ListDeviceCompliancePolicies(ctx, options)
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

	result = ListDeviceCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
