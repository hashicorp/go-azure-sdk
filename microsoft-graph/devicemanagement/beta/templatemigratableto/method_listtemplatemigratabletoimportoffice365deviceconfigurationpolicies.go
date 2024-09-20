package templatemigratableto

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

type ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntent
}

type ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntent
}

type ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions() ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions {
	return ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions{}
}

func (o ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateMigratableToImportOffice365DeviceConfigurationPolicies - Invoke action
// importOffice365DeviceConfigurationPolicies
func (c TemplateMigratableToClient) ListTemplateMigratableToImportOffice365DeviceConfigurationPolicies(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions) (result ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/migratableTo/importOffice365DeviceConfigurationPolicies", id.ID()),
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
		Values *[]beta.DeviceManagementIntent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesComplete retrieves all the results into a single object
func (c TemplateMigratableToClient) ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesComplete(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions) (ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteResult, error) {
	return c.ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteMatchingPredicate(ctx, id, options, DeviceManagementIntentOperationPredicate{})
}

// ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateMigratableToClient) ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesOperationOptions, predicate DeviceManagementIntentOperationPredicate) (result ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntent, 0)

	resp, err := c.ListTemplateMigratableToImportOffice365DeviceConfigurationPolicies(ctx, id, options)
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

	result = ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
