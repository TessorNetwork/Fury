package v3_0_0

import (
	"strings"

	v220did "github.com/tessornetwork/fury/x/did/legacy/v2.2.0"
	"github.com/tessornetwork/fury/x/did/types"
)

// Migrate accepts exported genesis state from v2.2.0 and migrates it to v3.0.0
func Migrate(oldGenState v220did.GenesisState) *types.GenesisState {

	identities := []*types.Identity{}

	// support data structure for no duplicates
	appears := map[string]struct{}{}

	for _, v220didDocument := range oldGenState.DidDocuments {
		if _, found := appears[string(v220didDocument.ID.String())]; !found {
			identity := fromDidDocumentToIdentity(v220didDocument)
			identities = append(identities, identity)
			appears[string(v220didDocument.ID.String())] = struct{}{}
		}
	}

	return &types.GenesisState{Identities: identities}
}

// TODO check for regex for pem
func publicKeyPemToMultiBase(pkPem string) (pkMultiBase string) {
	pkMultiBase = pkPem
	pkMultiBase = strings.ReplaceAll(pkMultiBase, "\n", "")
	pkMultiBase = strings.ReplaceAll(pkMultiBase, "\r", "")
	pkMultiBase = strings.ReplaceAll(pkMultiBase, "-", "")
	pkMultiBase = strings.TrimPrefix(pkMultiBase, "BEGIN PUBLIC KEY")
	pkMultiBase = strings.TrimSuffix(pkMultiBase, "END PUBLIC KEY")
	// add multibase code for base64 (rfc4648 no padding)
	pkMultiBase = "m" + pkMultiBase
	return
}

func convertPubKeys(pubKeys v220did.PubKeys) (verificationMethods []*types.VerificationMethod) {

	for _, pubKey := range pubKeys {

		verificationMethod := types.VerificationMethod{
			ID:                 pubKey.ID,
			Type:               pubKey.Type,
			Controller:         pubKey.Controller.String(),
			PublicKeyMultibase: publicKeyPemToMultiBase(pubKey.PublicKeyPem),
		}

		verificationMethods = append(verificationMethods, &verificationMethod)
	}

	return
}

func convertService(services220 v220did.Services) (services300 []*types.Service) {

	for _, service220 := range services220 {
		service300 := types.Service{
			ID:              service220.ID,
			Type:            service220.Type,
			ServiceEndpoint: service220.ServiceEndpoint,
		}
		services300 = append(services300, &service300)
	}

	return
}

func fromDidDocumentToIdentity(ddo v220did.DidDocument) (identity *types.Identity) {
	identity = &types.Identity{
		DidDocument: &types.DidDocument{
			Context:              []string{ddo.Context},
			ID:                   ddo.ID.String(),
			VerificationMethod:   convertPubKeys(ddo.PubKeys),
			Authentication:       []string{},
			AssertionMethod:      []string{},
			KeyAgreement:         []string{},
			CapabilityInvocation: []string{},
			CapabilityDelegation: []string{},
			Service:              convertService(ddo.Service),
		},
		Metadata: &types.Metadata{
			Created: ddo.Proof.Created.Format(types.ComplaintW3CTime),
			Updated: ddo.Proof.Created.Format(types.ComplaintW3CTime),
		},
	}

	return
}
