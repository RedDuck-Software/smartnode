package minipool

import (
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/types/api"
)


func getStatus(c *cli.Context) (*api.MinipoolStatusResponse, error) {

    // Get services
    if err := services.RequireNodeRegistered(c); err != nil { return nil, err }
    w, err := services.GetWallet(c)
    if err != nil { return nil, err }
    rp, err := services.GetRocketPool(c)
    if err != nil { return nil, err }

    // Response
    response := api.MinipoolStatusResponse{}

    // Get minipool details
    nodeAccount, _ := w.GetNodeAccount()
    details, err := getNodeMinipoolDetails(rp, nodeAccount.Address)
    if err != nil {
        return nil, err
    }
    response.Minipools = details

    // Return response
    return &response, nil

}
