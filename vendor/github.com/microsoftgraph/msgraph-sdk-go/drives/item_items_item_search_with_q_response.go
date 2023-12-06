package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemSearchWithQResponse 
// Deprecated: This class is obsolete. Use searchWithQGetResponse instead.
type ItemItemsItemSearchWithQResponse struct {
    ItemItemsItemSearchWithQGetResponse
}
// NewItemItemsItemSearchWithQResponse instantiates a new ItemItemsItemSearchWithQResponse and sets the default values.
func NewItemItemsItemSearchWithQResponse()(*ItemItemsItemSearchWithQResponse) {
    m := &ItemItemsItemSearchWithQResponse{
        ItemItemsItemSearchWithQGetResponse: *NewItemItemsItemSearchWithQGetResponse(),
    }
    return m
}
// CreateItemItemsItemSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemSearchWithQResponse(), nil
}
// ItemItemsItemSearchWithQResponseable 
// Deprecated: This class is obsolete. Use searchWithQGetResponse instead.
type ItemItemsItemSearchWithQResponseable interface {
    ItemItemsItemSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
