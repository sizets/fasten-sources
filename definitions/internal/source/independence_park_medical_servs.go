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

// https://fhir-myrecord.cerner.com/r4/fdbc9673-c618-4425-b458-9d4512be1bb3/metadata
func GetSourceIndependenceParkMedicalServs(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/fdbc9673-c618-4425-b458-9d4512be1bb3/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/fdbc9673-c618-4425-b458-9d4512be1bb3/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/fdbc9673-c618-4425-b458-9d4512be1bb3"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/fdbc9673-c618-4425-b458-9d4512be1bb3"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeIndependenceParkMedicalServs]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Independence Park Medical Servs."
	sourceDef.SourceType = pkg.SourceTypeIndependenceParkMedicalServs
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}