package v3_0_0

import (
	"reflect"
	"testing"
	"time"

	v220furykyc "github.com/tessornetwork/fury/x/furykyc/legacy/v2.2.0"
	"github.com/tessornetwork/fury/x/furykyc/types"
	ctypes "github.com/tessornetwork/fury/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	testUser01, _ = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")
	testUser02, _ = sdk.AccAddressFromBech32("cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh")
	testUser03, _ = sdk.AccAddressFromBech32("cosmos1h7tw92a66gr58pxgmf6cc336lgxadpjz5d5psf")
	testTsp, _    = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")
	timeNow       = time.Now()
	mExp01        = timeNow.Add(time.Hour * 24 * 60)
	mExp02        = timeNow.Add(time.Hour * 24 * 120)
	testDenom     = "ufury"
)

func TestMigrate(t *testing.T) {
	type args struct {
		v220GenState v220furykyc.GenesisState
	}
	tests := []struct {
		name string
		args args
		want *types.GenesisState
	}{
		{
			name: "empty",
			args: args{v220GenState: v220furykyc.GenesisState{}},
			want: &types.GenesisState{},
		},
		{
			name: "genesis state corectly migrated",
			args: args{
				v220GenState: v220furykyc.GenesisState{
					LiquidityPoolAmount: sdk.Coins{
						sdk.Coin{
							Denom:  testDenom,
							Amount: sdk.NewInt(1000000),
						},
					},
					Invites: v220furykyc.Invites{
						v220furykyc.Invite{
							Sender:           testUser01,
							SenderMembership: "black",
							Status:           v220furykyc.InviteStatusRewarded,
							User:             testUser02,
						},
						v220furykyc.Invite{
							Sender:           testUser02,
							SenderMembership: "black",
							Status:           v220furykyc.InviteStatusPending,
							User:             testUser03,
						},
					},
					Memberships: v220furykyc.Memberships{
						v220furykyc.Membership{
							Owner:          testUser02,
							TspAddress:     testTsp,
							MembershipType: "gold",
							ExpiryAt:       mExp01,
						},
						v220furykyc.Membership{
							Owner:          testUser03,
							TspAddress:     testTsp,
							MembershipType: "bronze",
							ExpiryAt:       mExp02,
						},
					},
					TrustedServiceProviders: ctypes.Addresses{
						testTsp,
					},
				},
			},
			want: &types.GenesisState{
				LiquidityPoolAmount: sdk.Coins{
					sdk.Coin{
						Denom:  testDenom,
						Amount: sdk.NewInt(1000000),
					},
				},
				Invites: []*types.Invite{
					&types.Invite{
						Sender:           testUser01.String(),
						SenderMembership: "black",
						Status:           uint64(1),
						User:             testUser02.String(),
					},
					&types.Invite{
						Sender:           testUser02.String(),
						SenderMembership: "black",
						Status:           uint64(0),
						User:             testUser03.String(),
					},
				},
				Memberships: []*types.Membership{
					&types.Membership{
						Owner:          testUser02.String(),
						TspAddress:     testTsp.String(),
						MembershipType: "gold",
						ExpiryAt:       &mExp01,
					},
					&types.Membership{
						Owner:          testUser03.String(),
						TspAddress:     testTsp.String(),
						MembershipType: "bronze",
						ExpiryAt:       &mExp02,
					},
				},
				TrustedServiceProviders: []string{
					testTsp.String(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Migrate(tt.args.v220GenState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
