const fs = require('fs');
const Web3 = require('web3');
const config = require('./config');

// Contracts
const RocketMinipoolSettings = JSON.parse(fs.readFileSync(config.rocketPoolPath + 'build/contracts/RocketMinipoolSettings.json'));

// Set minipool setting
async function setMinipoolSetting() {
    try {

        // Initialise web3
        const web3 = new Web3(config.providerUrl);

        // Parse arguments
        let args = process.argv.slice(2);
        if (args.length != 2) throw new Error('Usage: node set-minipool-setting.js setting value');
        let setting = args[0];
        let value = !!parseInt(args[1]);

        // Get network ID & accounts
        let networkId = await web3.eth.net.getId();
        let accounts = await web3.eth.getAccounts();

        // Initialise contracts
        let rocketMinipoolSettings = new web3.eth.Contract(RocketMinipoolSettings.abi, RocketMinipoolSettings.networks[networkId].address);

        // Set minipool setting
        await rocketMinipoolSettings.methods[setting](value).send({
            from: accounts[0],
            gas: 8000000,
        });

        // Log
        console.log('RocketMinipoolSettings setting ' + setting + ' successfully set to ' + value + '.');

    }
    catch (e) {
        console.log(e.message);
    }
}
setMinipoolSetting();