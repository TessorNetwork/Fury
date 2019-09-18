package id

import (
	"fmt"
	"strings"
	"testing"

	"github.com/commercionetwork/commercionetwork/x/id/internal/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var msgSetId = MsgSetIdentity{
	DidDocumentUri: keeper.TestDidDocumentUri,
	Owner:          keeper.TestOwnerAddress,
}

func TestValidMsg_StoreDoc(t *testing.T) {
	_, ctx, k := TestSetup()
	var handler = NewHandler(k)
	res := handler(ctx, msgSetId)
	require.True(t, res.IsOK())
}

func TestInvalidMsg(t *testing.T) {
	_, ctx, k := TestSetup()
	var handler = NewHandler(k)
	res := handler(ctx, sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, fmt.Sprintf("Unrecognized %s message type", ModuleName)))
}