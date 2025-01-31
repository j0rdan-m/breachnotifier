package config

import (
	"encoding/json"
	"os"
)

type State struct {
	Breaches map[string][]string `json:"breaches"` // Email -> Liste des sources
}

func LoadState(filename string) (*State, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		// Si le fichier n'existe pas, initialiser une carte vide
		return &State{Breaches: make(map[string][]string)}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state State
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		return nil, err
	}

	// Initialiser la carte si elle est nil
	if state.Breaches == nil {
		state.Breaches = make(map[string][]string)
	}

	return &state, nil
}

func SaveState(filename string, state *State) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pour des fichiers JSON lisibles
	return encoder.Encode(state)
}

func GetNewBreaches(email string, currentSources []string, state *State) []string {
	existingSources, exists := state.Breaches[email]
	if !exists {
		return currentSources // Si l'e-mail est nouveau, toutes les sources sont nouvelles
	}

	// Comparer les sources actuelles avec celles déjà enregistrées
	newSources := []string{}
	for _, source := range currentSources {
		found := false
		for _, existing := range existingSources {
			if source == existing {
				found = true
				break
			}
		}
		if !found {
			newSources = append(newSources, source)
		}
	}

	return newSources
}
