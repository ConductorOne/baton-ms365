package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryPurgeDataOperation 
type EdiscoveryPurgeDataOperation struct {
    CaseOperation
}
// NewEdiscoveryPurgeDataOperation instantiates a new ediscoveryPurgeDataOperation and sets the default values.
func NewEdiscoveryPurgeDataOperation()(*EdiscoveryPurgeDataOperation) {
    m := &EdiscoveryPurgeDataOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryPurgeDataOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryPurgeDataOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryPurgeDataOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryPurgeDataOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EdiscoveryPurgeDataOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// EdiscoveryPurgeDataOperationable 
type EdiscoveryPurgeDataOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
