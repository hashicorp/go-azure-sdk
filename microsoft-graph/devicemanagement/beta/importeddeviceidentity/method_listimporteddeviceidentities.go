package importeddeviceidentity

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

type ListImportedDeviceIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitiesOperationOptions struct {
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

func DefaultListImportedDeviceIdentitiesOperationOptions() ListImportedDeviceIdentitiesOperationOptions {
	return ListImportedDeviceIdentitiesOperationOptions{}
}

func (o ListImportedDeviceIdentitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListImportedDeviceIdentitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListImportedDeviceIdentitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListImportedDeviceIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedDeviceIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedDeviceIdentities - Get importedDeviceIdentities from deviceManagement. The imported device identities.
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentities(ctx context.Context, options ListImportedDeviceIdentitiesOperationOptions) (result ListImportedDeviceIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListImportedDeviceIdentitiesCustomPager{},
		Path:          "/deviceManagement/importedDeviceIdentities",
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

	temp := make([]beta.ImportedDeviceIdentity, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalImportedDeviceIdentityImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ImportedDeviceIdentity (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListImportedDeviceIdentitiesComplete retrieves all the results into a single object
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitiesComplete(ctx context.Context, options ListImportedDeviceIdentitiesOperationOptions) (ListImportedDeviceIdentitiesCompleteResult, error) {
	return c.ListImportedDeviceIdentitiesCompleteMatchingPredicate(ctx, options, ImportedDeviceIdentityOperationPredicate{})
}

// ListImportedDeviceIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitiesCompleteMatchingPredicate(ctx context.Context, options ListImportedDeviceIdentitiesOperationOptions, predicate ImportedDeviceIdentityOperationPredicate) (result ListImportedDeviceIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedDeviceIdentity, 0)

	resp, err := c.ListImportedDeviceIdentities(ctx, options)
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

	result = ListImportedDeviceIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
