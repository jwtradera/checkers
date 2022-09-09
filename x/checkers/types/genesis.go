
package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
    return &GenesisState{
        StoredGameList: []*StoredGame{},
        NextGame: &NextGame{
            Creator: "",
            IdValue: uint64(DefaultIndex),
        },
    }
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in storedGame
	storedGameIndexMap := make(map[string]bool)

	for _, elem := range gs.StoredGameList {
		if _, ok := storedGameIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[elem.Index] = true
	}

	return nil
}