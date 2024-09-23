package entitlementmanagementaccesspackagecatalogaccesspackagecustomworkflowextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CustomCalloutExtension
}

type ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CustomCalloutExtension
}

type ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions() ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions {
	return ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensions - List accessPackageCustomWorkflowExtensions.
// Get a list of the accessPackageAssignmentRequestWorkflowExtension and accessPackageAssignmentWorkflowExtension
// objects and their properties. The resulting list includes all the customAccessPackageWorkflowExtension objects for
// the catalog that the caller has access to read. Each object includes an @odata.type property that indicates whether
// the object is an accessPackageAssignmentRequestWorkflowExtension or an accessPackageAssignmentWorkflowExtension.
func (c EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient) ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensions(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions) (result ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageCustomWorkflowExtensions", id.ID()),
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

	temp := make([]beta.CustomCalloutExtension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalCustomCalloutExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.CustomCalloutExtension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient) ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions) (ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx, id, options, CustomCalloutExtensionOperationPredicate{})
}

// ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient) ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsOperationOptions, predicate CustomCalloutExtensionOperationPredicate) (result ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteResult, err error) {
	items := make([]beta.CustomCalloutExtension, 0)

	resp, err := c.ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensions(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
