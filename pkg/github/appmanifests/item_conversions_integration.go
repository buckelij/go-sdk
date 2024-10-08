package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemConversionsIntegration struct {
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integration
}
// NewItemConversionsIntegration instantiates a new ItemConversionsIntegration and sets the default values.
func NewItemConversionsIntegration()(*ItemConversionsIntegration) {
    m := &ItemConversionsIntegration{
        Integration: *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NewIntegration(),
    }
    return m
}
// CreateItemConversionsIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemConversionsIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsIntegration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemConversionsIntegration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Integration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ItemConversionsIntegration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Integration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ItemConversionsIntegrationable interface {
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
