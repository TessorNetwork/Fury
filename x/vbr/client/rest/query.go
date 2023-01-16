package rest

import (
	"fmt"
	"net/http"

	"github.com/tessornetwork/fury/x/vbr/types"
	"github.com/cosmos/cosmos-sdk/client"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/vbr/funds", getRetrieveBlockRewardsPoolFunds(cliCtx)).Methods("GET")
	r.HandleFunc("/vbr/params", getParamsHandler(cliCtx)).Methods("GET")
}


// ----------------------------------
// --- Vbr
// ----------------------------------

func getRetrieveBlockRewardsPoolFunds(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryBlockRewardsPoolFunds)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w,
				http.StatusInternalServerError,
				fmt.Sprintf("Could not get total funds amount: \n %s", err),
			)
		}

		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}


func getParamsHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}