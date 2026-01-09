package entitlementmanagementcatalogcustomworkflowextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomCalloutExtension
}

type ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomCalloutExtension
}

type ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions struct {
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

func DefaultListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions() ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions {
	return ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions{}
}

func (o ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementCatalogCustomWorkflowExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogCustomWorkflowExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogCustomWorkflowExtensions - List accessPackageCustomWorkflowExtensions. Get a list of
// the accessPackageAssignmentRequestWorkflowExtension and accessPackageAssignmentWorkflowExtension objects and their
// properties. The resulting list includes all the customAccessPackageWorkflowExtension objects for the catalog that the
// caller has access to read. Each object includes an @odata.type property that indicates whether the object is an
// accessPackageAssignmentRequestWorkflowExtension or an accessPackageAssignmentWorkflowExtension.
func (c EntitlementManagementCatalogCustomWorkflowExtensionClient) ListEntitlementManagementCatalogCustomWorkflowExtensions(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions) (result ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementCatalogCustomWorkflowExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/customWorkflowExtensions", id.ID()),
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

	temp := make([]stable.CustomCalloutExtension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalCustomCalloutExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.CustomCalloutExtension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListEntitlementManagementCatalogCustomWorkflowExtensionsComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogCustomWorkflowExtensionClient) ListEntitlementManagementCatalogCustomWorkflowExtensionsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions) (ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteResult, error) {
	return c.ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx, id, options, CustomCalloutExtensionOperationPredicate{})
}

// ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogCustomWorkflowExtensionClient) ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogId, options ListEntitlementManagementCatalogCustomWorkflowExtensionsOperationOptions, predicate CustomCalloutExtensionOperationPredicate) (result ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteResult, err error) {
	items := make([]stable.CustomCalloutExtension, 0)

	resp, err := c.ListEntitlementManagementCatalogCustomWorkflowExtensions(ctx, id, options)
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

	result = ListEntitlementManagementCatalogCustomWorkflowExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
