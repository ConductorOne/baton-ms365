package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContactsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactsDeltaResponse struct {
    ItemContactsDeltaGetResponse
}
// NewItemContactsDeltaResponse instantiates a new ItemContactsDeltaResponse and sets the default values.
func NewItemContactsDeltaResponse()(*ItemContactsDeltaResponse) {
    m := &ItemContactsDeltaResponse{
        ItemContactsDeltaGetResponse: *NewItemContactsDeltaGetResponse(),
    }
    return m
}
// CreateItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactsDeltaResponse(), nil
}
// ItemContactsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactsDeltaResponseable interface {
    ItemContactsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}