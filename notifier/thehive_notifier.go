package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TheHiveNotifier struct {
	ApiURL       string
	ApiKey       string
	Organisation string
}

func NewTheHiveNotifier(apiURL, apiKey, organisation string) *TheHiveNotifier {
	return &TheHiveNotifier{
		ApiURL:       apiURL,
		ApiKey:       apiKey,
		Organisation: organisation,
	}
}

func (n *TheHiveNotifier) Notify(entry NotificationEntry) error {
	// Construire la liste des observables
	observables := []map[string]interface{}{
		{"dataType": "mail", "data": entry.Email}, // Observable pour l'e-mail
	}

	// Ajouter les champs compromis en tant qu'observables
	for _, field := range entry.Fields {
		observables = append(observables, map[string]interface{}{
			"dataType": "text",
			"data":     fmt.Sprintf("Compromised field: %s", field),
		})
	}

	// Ajouter les sources de compromissions en tant qu'observables
	for _, source := range entry.Sources { // entry.Sources est déjà de type []string
		observables = append(observables, map[string]interface{}{
			"dataType": "text",
			"data":     fmt.Sprintf("Source: %s", source),
		})
	}

	// Construire l'alerte pour TheHive
	alert := map[string]interface{}{
		"type":        "external",                                                // Type d'alerte
		"source":      "family",                                                  // Source (personnalisable si nécessaire)
		"severity":    2,                                                         // Niveau de sévérité (ajustable)
		"sourceRef":   fmt.Sprintf("%s-%d", entry.Email, time.Now().UnixMilli()), // Référence unique
		"title":       fmt.Sprintf("New Breach Alert: %s", entry.Email),
		"description": fmt.Sprintf("Detected %d breaches for %s.\nFields: %v\nSources: %v", entry.Found, entry.Email, entry.Fields, entry.Sources),
		"date":        time.Now().UnixMilli(), // Timestamp en millisecondes
		"observables": observables,
	}

	// Encoder les données en JSON
	payload, err := json.Marshal(alert)
	if err != nil {
		return fmt.Errorf("failed to marshal alert: %w", err)
	}

	// Construire l'URL complète pour l'endpoint
	url := fmt.Sprintf("%s/api/v1/alert", n.ApiURL)

	// Créer une requête HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Ajouter les en-têtes requis
	req.Header.Set("Authorization", "Bearer "+n.ApiKey)
	req.Header.Set("X-Organisation", n.Organisation)
	req.Header.Set("Content-Type", "application/json")

	// Envoyer la requête
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Lire et afficher la réponse en cas d'erreur
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send notification to TheHive, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
