
# BreachNotifier

**BreachNotifier** is a tool designed to monitor data breaches associated with a list of email addresses. It leverages services like **LeakCheck** to verify data leaks, can log results for SIEM tools (e.g., Wazuh), and sends notifications to systems like **TheHive**. This project is built to be extensible, modular, and easily configurable.

---

## 🚀 Features
- Automatically checks for email data breaches using the LeakCheck API.
- Modular architecture: interchangeable loggers and notifiers.
- Sends notifications to **TheHive** with detailed breach alerts.
- Logs results for SIEM tools like **Wazuh**.
- Tracks breach state to only detect new compromises.
- Simple configuration via a YAML file.

---

## 📦 Installation

### 1. Prerequisites
- [Go](https://go.dev/) (version 1.18 or later)
- An account or API key for the **LeakCheck** service.
- A **TheHive** instance (optional for notifications).
- (Optional) A SIEM tool like **Wazuh** for logging.

### 2. Clone the repository
```bash
git clone https://github.com/<your-username>/breachnotifier.git
cd breachnotifier
```

### 3. Build the binary
```bash
go build -o breachnotifier
```

### 4. Run the installation script
```bash
sudo ./deploy/install.sh
```

---

## ⚙️ Configuration

Edit the `config.yaml` file to customize the behavior of BreachNotifier:

```yaml
emails:
  - "test@example.com"
  - "hello@world.com"
logger:
  type: "wazuh"             # Logger type (e.g., wazuh, elk). Leave empty to disable.
  file_path: "logs.json"    # Path to the log file.
checker:
  type: "leakcheck"         # Service used to check for breaches (currently LeakCheck).
notifier:
  type: "thehive"           # Notifier type (e.g., thehive). Leave empty to disable.
  api_url: "http://localhost:9000"
  api_key: "your_api_key"
  organisation: "family"    # TheHive organisation
```

- **`emails`**: List of email addresses to monitor.
- **`logger`**: Logger configuration (optional).
- **`checker`**: Specifies the service used for breach checks.
- **`notifier`**: Configuration for sending notifications via a service like TheHive (optional).

---

## 🛠️ Usage

### Run the program
```bash
./breachnotifier
```

### Verify the service
Once installed, you can verify the systemd timer:
```bash
systemctl status breachnotifier.timer
```

---

## 🧑‍💻 Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a branch with a descriptive name:
   ```bash
   git checkout -b feature/my-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add my new feature"
   ```
4. Push the branch:
   ```bash
   git push origin feature/my-feature
   ```
5. Open a pull request.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 💡 Future Features
- Integration with other breach checkers (e.g., HaveIBeenPwned).
- Additional notifiers like Splunk or PagerDuty.
- Reporting and exportable breach statistics.

---

## 📞 Support

For questions or issues, open an issue on GitHub or contact [your_email@example.com].

---

## 🌟 Acknowledgments
Thank you to everyone contributing to improving user security through open-source projects!
