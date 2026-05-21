# cronmap

Visualizer that parses crontab entries and renders a human-readable weekly schedule.

---

## Installation

```bash
go install github.com/yourusername/cronmap@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/cronmap.git && cd cronmap && go build ./...
```

---

## Usage

Pipe your crontab directly into `cronmap`:

```bash
crontab -l | cronmap
```

Or pass a crontab file as an argument:

```bash
cronmap -f /etc/cron.d/myjobs
```

**Example output:**

```
Mon  02:30  backup-db
Mon  08:00  send-report
Wed  02:30  backup-db
Fri  18:00  cleanup-logs
Sun  00:00  weekly-digest
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-f` | Path to a crontab file | stdin |
| `-tz` | Timezone for rendering | UTC |
| `--json` | Output schedule as JSON | false |

---

## Requirements

- Go 1.21+

---

## License

[MIT](LICENSE)