
package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

type Asset struct {
    DealerID    string `json:"dealerID"`
    MSISDN      string `json:"msisdn"`
    MPIN        string `json:"mpin"`
    Balance     string `json:"balance"`
    Status      string `json:"status"`
    TransAmount string `json:"transAmount"`
    TransType   string `json:"transType"`
    Remarks     string `json:"remarks"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    return nil
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin, balance, status, transAmount, transType, remarks string) error {
    asset := Asset{
        DealerID:    dealerID,
        MSISDN:      msisdn,
        MPIN:        mpin,
        Balance:     balance,
        Status:      status,
        TransAmount: transAmount,
        TransType:   transType,
        Remarks:     remarks,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(msisdn, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, msisdn string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(msisdn)
    if err != nil {
        return nil, err
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset %s does not exist", msisdn)
    }
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }
    return &asset, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, msisdn, balance, status string) error {
    asset, err := s.ReadAsset(ctx, msisdn)
    if err != nil {
        return err
    }
    asset.Balance = balance
    asset.Status = status

    updatedJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(msisdn, updatedJSON)
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        panic(fmt.Sprintf("Error creating chaincode: %v", err))
    }
    if err := chaincode.Start(); err != nil {
        panic(fmt.Sprintf("Error starting chaincode: %v", err))
    }
}
