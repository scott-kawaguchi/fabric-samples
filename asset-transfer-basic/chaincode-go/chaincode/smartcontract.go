package chaincode

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	votes := []VoteData{
		{
			User: User{
			ID: "user1",
			FirstName: "Jimmy",
			LastName: "Garoppolo",
			DriversLicenseID: "B2357345",
			ImageURIs: []string{"http://securevote.com/bucketid/image"},
			VideoURIs: []string{"http://securevote.com/bucketid/myvid"},
			SoundURIs: []string{"http://securevote.com/bucketid/mysoundbite"},
		    },
			Vote: "Jeff",
			Gateway: Gateway{
				ID: "gateway1",
				Location: "California coordinates: 36.778259, -119.417931",
			},
		},
		{
			User: User{
				ID: "user2",
				FirstName: "Stephen",
				LastName: "Curry",
				DriversLicenseID: "B2456745",
				ImageURIs: []string{"http://securevote.com/bucketid/image"},
				VideoURIs: []string{"http://securevote.com/bucketid/myvid"},
				SoundURIs: []string{"http://securevote.com/bucketid/mysoundbite"},
			},
			Vote: "Jeff",
			Gateway: Gateway{
				ID: "gateway1",
				Location: "California coordinates: 36.778259, -119.417931",
			},
		},
		{
			User: User{
				ID: "user3",
				FirstName: "Jeff",
				LastName: "Bezos",
				DriversLicenseID: "B2456745",
				ImageURIs: []string{"http://securevote.com/bucketid/image"},
				VideoURIs: []string{"http://securevote.com/bucketid/myvid"},
				SoundURIs: []string{"http://securevote.com/bucketid/mysoundbite"},
			},
			Vote: "Jeff",
			Gateway: Gateway{
				ID: "gateway2",
				Location: "Washington coordinates: 47.751076, -120.740135",
			},
		},

	}

	for _, vote := range votes {
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(vote.User.ID, voteJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateVote(ctx contractapi.TransactionContextInterface, id string, firstName string, lastName string,
	driversLicenseId string, imageUrls string, videoUrls string, soundUrls string, vote string, gatewayId string,
	gatewayLocation string) error {
	exists, err := s.VoteExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the user's vote %s already exists", id)
	}

	newvote := VoteData{
		User: User{
			ID: id,
			FirstName: firstName,
			LastName: lastName,
			DriversLicenseID: driversLicenseId,
			ImageURIs: strings.Split(imageUrls, ","),
			VideoURIs: strings.Split(videoUrls, ","),
			SoundURIs: strings.Split(soundUrls, ","),
		},
		Vote: vote,
		Gateway: Gateway{
			ID: gatewayId,
			Location: gatewayLocation,
		},
	}
	voteJSON, err := json.Marshal(newvote)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, voteJSON)
}

// ReadVote returns the voting data stored in the world state with given id.
func (s *SmartContract) ReadVote(ctx contractapi.TransactionContextInterface, id string) (*VoteData, error) {
	voteJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if voteJSON == nil {
		return nil, fmt.Errorf(" asset %s does not exist", id)
	}

	var voteData VoteData
	err = json.Unmarshal(voteJSON, &voteData)
	if err != nil {
		return nil, err
	}

	return &voteData, nil
}

//// UpdateAsset updates an existing asset in the world state with provided parameters.
//func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
//	exists, err := s.AssetExists(ctx, id)
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return fmt.Errorf("the asset %s does not exist", id)
//	}
//
//	// overwriting original asset with new asset
//	asset := Asset{
//		ID:             id,
//		Color:          color,
//		Size:           size,
//		Owner:          owner,
//		AppraisedValue: appraisedValue,
//	}
//	assetJSON, err := json.Marshal(asset)
//	if err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(id, assetJSON)
//}

//// DeleteAsset deletes an given asset from the world state.
//func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
//	exists, err := s.AssetExists(ctx, id)
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return fmt.Errorf("the asset %s does not exist", id)
//	}
//
//	return ctx.GetStub().DelState(id)
//}

// VoteExists returns true when voting data with given ID exists in world state
func (s *SmartContract) VoteExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	voteJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return voteJSON != nil, nil
}

//// TransferAsset updates the owner field of asset with given id in world state.
//func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
//	asset, err := s.ReadAsset(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	asset.Owner = newOwner
//	assetJSON, err := json.Marshal(asset)
//	if err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(id, assetJSON)
//}

// GetAllVotes returns all voting data found in world state
func (s *SmartContract) GetAllVotes(ctx contractapi.TransactionContextInterface) ([]*VoteData, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var votes []*VoteData
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var vote VoteData
		err = json.Unmarshal(queryResponse.Value, &vote)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &vote)
	}

	return votes, nil
}
