
const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

const app = express();
app.use(express.json());

app.post('/createAsset', async (req, res) => {
    try {
        const ccpPath = path.resolve(__dirname, 'connection.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
        const walletPath = path.join(__dirname, 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: 'appUser',
            discovery: { enabled: true, asLocalhost: true }
        });

        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('assetContract');

        const { dealerID, msisdn, mpin, balance, status, transAmount, transType, remarks } = req.body;
        await contract.submitTransaction('CreateAsset', dealerID, msisdn, mpin, balance, status, transAmount, transType, remarks);

        await gateway.disconnect();
        res.send('Asset created successfully');
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(error.toString());
    }
});

app.listen(3000, () => console.log('API server running on port 3000'));
