package watchtower

import (
    "log"
    "time"

    "github.com/rocket-pool/rocketpool-go/rocketpool"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/services/accounts"
)


// Settings
var processWithdrawalsInterval, _ = time.ParseDuration("1m")


// Start process withdrawals task
func startProcessWithdrawals(c *cli.Context) error {

    // Get services
    if err := services.WaitNodeRegistered(c, true); err != nil { return err }
    am, err := services.GetAccountManager(c)
    if err != nil { return err }
    rp, err := services.GetRocketPool(c)
    if err != nil { return err }

    // Process withdrawals at interval
    go (func() {
        for {
            if err := processWithdrawals(c, am, rp); err != nil {
                log.Println(err)
            }
            time.Sleep(processWithdrawalsInterval)
        }
    })()

    // Return
    return nil

}


// Process withdrawals
func processWithdrawals(c *cli.Context, am *accounts.AccountManager, rp *rocketpool.RocketPool) error {

    // Process withdrawals
    // TODO: implement

    // Return
    return nil

}
