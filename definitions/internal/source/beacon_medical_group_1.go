// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir-myrecord.cerner.com/r4/acdd0d06-725b-4e78-bc4c-92f3f27f3a8e/metadata
func GetSourceBeaconMedicalGroup1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/acdd0d06-725b-4e78-bc4c-92f3f27f3a8e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/acdd0d06-725b-4e78-bc4c-92f3f27f3a8e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/acdd0d06-725b-4e78-bc4c-92f3f27f3a8e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/acdd0d06-725b-4e78-bc4c-92f3f27f3a8e"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBeaconMedicalGroup1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Beacon Medical Group"
	sourceDef.SourceType = pkg.SourceTypeBeaconMedicalGroup1
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}