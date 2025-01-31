#!/bin/bash

set -e

echo "Installation de BreachNotifier..."

# Variables
BINARY_PATH="/usr/local/bin/breachnotifier"
WORKING_DIR="/opt/breachnotifier"
SERVICE_FILE="/etc/systemd/system/breachnotifier.service"
TIMER_FILE="/etc/systemd/system/breachnotifier.timer"

# Vérifier que le binaire existe
if [ ! -f "./breachnotifier" ]; then
    echo "Erreur : le fichier binaire 'breachnotifier' est introuvable. Compilez le projet avant d'exécuter ce script."
    exit 1
fi

# Créer le répertoire de travail
echo "Création du répertoire de travail : $WORKING_DIR"
sudo mkdir -p "$WORKING_DIR"
sudo cp config.yaml "$WORKING_DIR"
sudo cp -r db "$WORKING_DIR"

# Déplacer le binaire
echo "Déplacement du binaire vers $BINARY_PATH"
sudo cp breachnotifier "$BINARY_PATH"
sudo chmod +x "$BINARY_PATH"

# Créer un utilisateur dédié
echo "Création de l'utilisateur 'breachnotifier'"
if ! id -u breachnotifier >/dev/null 2>&1; then
    sudo useradd -r -d "$WORKING_DIR" -s /usr/sbin/nologin breachnotifier
fi

# Donner les droits à l'utilisateur dédié
sudo chown -R breachnotifier:breachnotifier "$WORKING_DIR"

# Installer le fichier de service
echo "Installation du service systemd"
sudo cp deploy/breachnotifier.service "$SERVICE_FILE"
sudo cp deploy/breachnotifier.timer "$TIMER_FILE"

# Recharger systemd et activer le timer
echo "Rechargement de systemd et activation du timer"
sudo systemctl daemon-reload
sudo systemctl enable breachnotifier.timer
sudo systemctl start breachnotifier.timer

echo "Installation terminée avec succès."
