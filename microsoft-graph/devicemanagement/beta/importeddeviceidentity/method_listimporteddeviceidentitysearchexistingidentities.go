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

type ListImportedDeviceIdentitySearchExistingIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions() ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions {
	return ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions{}
}

func (o ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedDeviceIdentitySearchExistingIdentities - Invoke action searchExistingIdentities
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentities(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest, options ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions) (result ListImportedDeviceIdentitySearchExistingIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager{},
		Path:          "/deviceManagement/importedDeviceIdentities/searchExistingIdentities",
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

// ListImportedDeviceIdentitySearchExistingIdentitiesComplete retrieves all the results into a single object
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentitiesComplete(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest, options ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions) (ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult, error) {
	return c.ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate(ctx, input, options, ImportedDeviceIdentityOperationPredicate{})
}

// ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest, options ListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions, predicate ImportedDeviceIdentityOperationPredicate) (result ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedDeviceIdentity, 0)

	resp, err := c.ListImportedDeviceIdentitySearchExistingIdentities(ctx, input, options)
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

	result = ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
