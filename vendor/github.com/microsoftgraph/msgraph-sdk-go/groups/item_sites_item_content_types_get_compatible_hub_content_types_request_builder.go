package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters invoke function getCompatibleHubContentTypes
type ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new GetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/getCompatibleHubContentTypes(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder instantiates a new GetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCompatibleHubContentTypes
// Deprecated: This method is obsolete. Use GetAsGetCompatibleHubContentTypesGetResponse instead.
func (m *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemContentTypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentTypesGetCompatibleHubContentTypesResponseable), nil
}
// GetAsGetCompatibleHubContentTypesGetResponse invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) GetAsGetCompatibleHubContentTypesGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemContentTypesGetCompatibleHubContentTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentTypesGetCompatibleHubContentTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentTypesGetCompatibleHubContentTypesGetResponseable), nil
}
// ToGetRequestInformation invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemSitesItemContentTypesGetCompatibleHubContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
