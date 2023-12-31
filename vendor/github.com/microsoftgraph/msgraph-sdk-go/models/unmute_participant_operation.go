package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmuteParticipantOperation 
type UnmuteParticipantOperation struct {
    CommsOperation
}
// NewUnmuteParticipantOperation instantiates a new unmuteParticipantOperation and sets the default values.
func NewUnmuteParticipantOperation()(*UnmuteParticipantOperation) {
    m := &UnmuteParticipantOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateUnmuteParticipantOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmuteParticipantOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnmuteParticipantOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmuteParticipantOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UnmuteParticipantOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// UnmuteParticipantOperationable 
type UnmuteParticipantOperationable interface {
    CommsOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
