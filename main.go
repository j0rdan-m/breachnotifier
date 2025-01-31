package main

import (
	"fmt"
	"time"

	"breachnotifier/checker"
	"breachnotifier/config"
	"breachnotifier/logger"
	"breachnotifier/notifier"
)

func main() {
	// Charger la configuration
	conf, stateFile, state, log, chk, notif, shouldReturn := initBreachNotifier()
	if shouldReturn {
		return
	}

	// Vérifier les e-mails
	for _, email := range conf.Emails {
		fmt.Printf("Vérification de l'email : %s\n", email)

		// Respecter la limite d'une requête par seconde
		time.Sleep(1 * time.Second)

		result, err := chk.CheckEmail(email)
		if err != nil {
			fmt.Printf("Erreur pour %s : %v\n", email, err)
			continue
		}

		// Obtenir les nouvelles breaches uniquement
		newSources := filterOnNewBreaches(result, email, state)
		if len(newSources) == 0 {
			fmt.Printf("Aucune nouvelle compromission détectée pour %s\n", email)
			continue
		}

		fmt.Printf("Nouvelles compromissions détectées pour %s : %v\n", email, newSources)

		// Mettre à jour l'état
		state.Breaches[email] = append(state.Breaches[email], newSources...)

		// Envoyer une notification si le notifier est défini
		notifyNewBreaches(notif, email, newSources, result)

		// Écrire dans les logs si le logger est défini
		logNewBreaches(log, email, newSources, result)
	}

	// Sauvegarder l'état mis à jour
	if err := config.SaveState(stateFile, state); err != nil {
		fmt.Println("Erreur lors de la sauvegarde de l'état :", err)
	}
}

func logNewBreaches(log logger.Logger, email string, newSources []string, result *checker.Response) {
	if log != nil {
		logEntry := logger.LogEntry{
			Email:     email,
			Found:     len(newSources),
			Fields:    result.Fields,
			Sources:   newSources,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		if err := log.WriteLog(logEntry); err != nil {
			fmt.Printf("Erreur lors de l'écriture du log pour %s : %v\n", email, err)
		}
	}
}

func notifyNewBreaches(notif notifier.Notifier, email string, newSources []string, result *checker.Response) {
	if notif != nil {
		notifEntry := notifier.NotificationEntry{
			Email:     email,
			Found:     len(newSources),
			Fields:    result.Fields,
			Sources:   newSources,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		if err := notif.Notify(notifEntry); err != nil {
			fmt.Printf("Erreur lors de l'envoi de la notification pour %s : %v\n", email, err)
		}
	}
}

func filterOnNewBreaches(result *checker.Response, email string, state *config.State) []string {
	currentSources := []string{}
	for _, source := range result.Sources {
		currentSources = append(currentSources, source.Name)
	}
	newSources := config.GetNewBreaches(email, currentSources, state)

	return newSources
}

func initBreachNotifier() (*config.Config, string, *config.State, logger.Logger, checker.Checker, notifier.Notifier, bool) {
	configFile := "config.yaml"
	conf, err := config.LoadConfig(configFile)
	if err != nil {
		fmt.Println("Erreur lors du chargement de la configuration :", err)
		return nil, "",

			// Charger l'état
			nil, nil, nil, nil, true
	}

	stateFile := "db/state.json"
	state, shouldReturn := loadState(stateFile)
	if shouldReturn {
		return nil, "",

			// Initialiser le logger (optionnel)
			nil, nil, nil, nil, true
	}

	log, shouldReturn := loadLogger(conf)
	if shouldReturn {
		return nil, "",

			// Initialiser le checker (obligatoire)
			nil, nil, nil, nil, true
	}

	chk, err := checker.GetChecker(conf.Checker.Type)
	if err != nil {
		fmt.Println("Erreur lors de l'initialisation du checker :", err)
		return nil, "",

			// Initialiser le notifier (optionnel)
			nil, nil, nil, nil, true
	}

	var notif notifier.Notifier
	notif, shouldReturn1 := initNotifier(conf, notif)
	if shouldReturn1 {
		return nil, "", nil, nil, nil, nil, true
	}
	return conf, stateFile, state, log, chk, notif, shouldReturn
}

func initNotifier(conf *config.Config, notif notifier.Notifier) (notifier.Notifier, bool) {
	if conf.Notifier != nil {
		var notifErr error
		notif, notifErr = notifier.GetNotifier(
			conf.Notifier.Type,
			conf.Notifier.ApiURL,
			conf.Notifier.ApiKey,
			conf.Notifier.Organisation,
		)
		if notifErr != nil {
			fmt.Println("Erreur lors de l'initialisation du notifier :", notifErr)
			return nil, true
		}
	} else {
		fmt.Println("Notifier désactivé.")
	}
	return notif, false
}

func loadState(stateFile string) (*config.State, bool) {
	state, err := config.LoadState(stateFile)
	if err != nil {
		fmt.Println("Erreur lors du chargement de l'état :", err)
		return nil, true
	}
	return state, false
}

func loadLogger(conf *config.Config) (logger.Logger, bool) {
	var log logger.Logger
	if conf.Logger != nil {
		var err error
		log, err = logger.GetLogger(conf.Logger.Type, conf.Logger.FilePath)
		if err != nil {
			fmt.Println("Erreur lors de l'initialisation du logger :", err)
			return nil, true
		}
	} else {
		fmt.Println("Logger désactivé.")
	}
	return log, false
}
