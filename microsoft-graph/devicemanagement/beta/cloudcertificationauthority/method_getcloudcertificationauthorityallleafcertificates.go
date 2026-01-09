package cloudcertificationauthority

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCloudCertificationAuthorityAllLeafCertificatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudCertificationAuthorityLeafCertificate
}

type GetCloudCertificationAuthorityAllLeafCertificatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudCertificationAuthorityLeafCertificate
}

type GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetCloudCertificationAuthorityAllLeafCertificatesOperationOptions() GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions {
	return GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions{}
}

func (o GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions) ToOData() *odata.Query {
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

func (o GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCloudCertificationAuthorityAllLeafCertificatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCloudCertificationAuthorityAllLeafCertificatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCloudCertificationAuthorityAllLeafCertificates - Invoke action getAllCloudCertificationAuthorityLeafCertificates
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllLeafCertificates(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions) (result GetCloudCertificationAuthorityAllLeafCertificatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCloudCertificationAuthorityAllLeafCertificatesCustomPager{},
		Path:          fmt.Sprintf("%s/getAllCloudCertificationAuthorityLeafCertificates", id.ID()),
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
		Values *[]beta.CloudCertificationAuthorityLeafCertificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetCloudCertificationAuthorityAllLeafCertificatesComplete retrieves all the results into a single object
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllLeafCertificatesComplete(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions) (GetCloudCertificationAuthorityAllLeafCertificatesCompleteResult, error) {
	return c.GetCloudCertificationAuthorityAllLeafCertificatesCompleteMatchingPredicate(ctx, id, options, CloudCertificationAuthorityLeafCertificateOperationPredicate{})
}

// GetCloudCertificationAuthorityAllLeafCertificatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllLeafCertificatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllLeafCertificatesOperationOptions, predicate CloudCertificationAuthorityLeafCertificateOperationPredicate) (result GetCloudCertificationAuthorityAllLeafCertificatesCompleteResult, err error) {
	items := make([]beta.CloudCertificationAuthorityLeafCertificate, 0)

	resp, err := c.GetCloudCertificationAuthorityAllLeafCertificates(ctx, id, options)
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

	result = GetCloudCertificationAuthorityAllLeafCertificatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
