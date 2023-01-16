package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

// package initialization for correct validation of tessornetwork addresses
func init() {
	configTestPrefixes()
}

func configTestPrefixes() {
	AccountAddressPrefix := "did:fury:"
	AccountPubKeyPrefix := AccountAddressPrefix + "pub"
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.Seal()
}

func TestService_isValid(t *testing.T) {

	tests := []struct {
		name    string
		service func() *Service
		wantErr bool
	}{
		{
			"valid",
			func() *Service {
				return &validServiceBar
			},
			false,
		},
		{
			"not defined",
			func() *Service {
				return nil
			},
			true,
		},
		{
			"{ID} empty",
			func() *Service {
				service := validServiceBar
				service.ID = ""
				return &service
			},
			true,
		},
		{
			"{ID} against the rules of RFC3986",
			func() *Service {
				service := validServiceBar
				service.ID = "$" + validServiceBar.ID
				return &service
			},
			true,
		},
		{
			"{ID} size too long",
			func() *Service {
				service := validServiceBar
				service.ID = service.ID + strings.Repeat("c", 1+serviceLenghtLimitID)
				return &service
			},
			true,
		},
		{
			"{type} empty",
			func() *Service {
				service := validServiceBar
				service.Type = ""
				return &service
			},
			true,
		},
		{
			"{type} size too long",
			func() *Service {
				service := validServiceBar
				service.Type = strings.Repeat("c", 1+serviceLenghtLimitType)
				return &service
			},
			true,
		},
		{
			"{serviceEndpoint} empty",
			func() *Service {
				service := validServiceBar
				service.ServiceEndpoint = ""
				return &service
			},
			true,
		},
		{
			"{serviceEndpoint} against the rules of RFC3986",
			func() *Service {
				service := validServiceBar
				service.ServiceEndpoint = "$" + validServiceBar.ServiceEndpoint
				return &service
			},
			true,
		},
		{
			"{serviceEndpoint} size too long",
			func() *Service {
				service := validServiceBar
				service.ServiceEndpoint = strings.Repeat("c", 1+serviceLenghtLimitServiceEndpoint)
				return &service
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.service().Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Service.isValid() for service %s error = %v, wantErr %v", tt.service(), err, tt.wantErr)
			}
		})
	}
}

func TestVerificationMethod_isValid(t *testing.T) {

	tests := []struct {
		name               string
		verificationMethod func() *VerificationMethod
		wantErr            bool
	}{
		{
			"valid",
			func() *VerificationMethod {
				return &validVerificationMethodRsaVerificationKey2018
			},
			false,
		},
		{
			"valid with RSA key with different decoding size should be valid",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64Padded) + validBase64RsaVerificationKey2018DecondingSize,
				}
			},
			false,
		},
		{
			"valid with RSA key with 2048 bytes should be valid",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64Padded) + validBase64RsaVerificationKey20182048,
				}
			},
			false,
		},
		{
			"valid with RSA key with 4096 bytes should be valid",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64Padded) + validBase64RsaVerificationKey20184096,
				}
			},
			false,
		},
		{
			"valid with RSA key OpenSSL-generated should be valid",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64Padded) + validBase64RsaVerificationKey2018OpenSSL,
				}
			},
			false,
		},
		{
			"valid padded RSA key with no padding multibase code should error",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64NoPadding) + validBase64RsaVerificationKey2018Padded,
				}
			},
			true,
		},
		{
			"invalid multibase code should error",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: "u" + validBase64RsaVerificationKey20182048,
				}
			},
			true,
		},
		{
			"valid no padding RSA key with padding multibase code should error",
			func() *VerificationMethod {
				return &VerificationMethod{
					ID:                 validDidSubject + RsaVerificationKey2018NameSuffix,
					Type:               RsaVerificationKey2018,
					Controller:         validDidSubject,
					PublicKeyMultibase: string(MultibaseCodeBase64Padded) + validBase64RsaVerificationKey2018NoPadding,
				}
			},
			true,
		},
		{
			"not defined",
			func() *VerificationMethod {
				return nil
			},
			true,
		},
		{
			"{ID} empty",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.ID = ""
				return &verificationMethod
			},
			true,
		},
		{
			"{ID} against the DID url specification",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.ID = "$" + validVerificationMethodRsaVerificationKey2018.ID
				return &verificationMethod
			},
			true,
		},
		{
			"{type} empty",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.Type = ""
				return &verificationMethod
			},
			true,
		},
		{
			"{type} not supported",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.Type = "NotSupported2077"
				return &verificationMethod
			},
			true,
		},
		{
			"{type} and {ID} mismatch for " + RsaVerificationKey2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.ID = validDidSubject + RsaSignature2018NameSuffix
				return &verificationMethod
			},
			true,
		},
		{
			"{type} and {ID} mismatch for " + RsaSignature2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaSignature2018
				verificationMethod.ID = validDidSubject + RsaVerificationKey2018NameSuffix
				return &verificationMethod
			},
			true,
		},
		{
			"{controller} empty",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.Controller = ""
				return &verificationMethod
			},
			true,
		},
		{
			"{controller} against the DID specification",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.Controller = "$" + validDidSubject
				return &verificationMethod
			},
			true,
		},
		{
			"{controller} different from subject",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.Controller = validDidNoSubject
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} empty",
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.PublicKeyMultibase = ""
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid format for {type} " + RsaVerificationKey2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.PublicKeyMultibase = verificationMethod.PublicKeyMultibase[1:]
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid format for {type} " + RsaSignature2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaSignature2018
				verificationMethod.PublicKeyMultibase = verificationMethod.PublicKeyMultibase[1:]
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid base64 encoding for {type} " + RsaVerificationKey2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.PublicKeyMultibase = verificationMethod.PublicKeyMultibase + "-"
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid base64 encoding for {type} " + RsaSignature2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaSignature2018
				verificationMethod.PublicKeyMultibase = verificationMethod.PublicKeyMultibase + "-"
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid key for {type} " + RsaVerificationKey2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaVerificationKey2018
				verificationMethod.PublicKeyMultibase = string(MultibaseCodeBase64Padded) + invalidBase64RSAKey
				return &verificationMethod
			},
			true,
		},
		{
			"{publicKeyMultibase} invalid key for {type} " + RsaSignature2018,
			func() *VerificationMethod {
				verificationMethod := validVerificationMethodRsaSignature2018
				verificationMethod.PublicKeyMultibase = string(MultibaseCodeBase64Padded) + invalidBase64RSAKey
				return &verificationMethod
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEqual(t, validDidSubject, validDidNoSubject)
			if err := tt.verificationMethod().Validate(validDidSubject); (err != nil) != tt.wantErr {
				t.Errorf("VerificationMethod.isValid() for verificationMethod %s error = %v, wantErr %v", tt.verificationMethod(), err, tt.wantErr)
			}
		})
	}
}

func Test_isValidDidCom(t *testing.T) {

	tests := []struct {
		name    string
		did     string
		wantErr bool
	}{
		{"didSubject", validDidSubject, false},
		{"didNoSubject", validDidNoSubject, false},
		{"not valid suffix", validDidSubject + "$", true},
		{"not valid prefix", "$" + validDidSubject, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.did); (err != nil) != tt.wantErr {
				t.Errorf("isValidDidCom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
