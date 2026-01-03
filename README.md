# Dawn

A CLI tool for tracking morning routines. Visualize your consistency with streaks.

## Installation

```bash
go install github.com/matsuhaya/dawn@latest
```

## Usage

### Record a routine

Record what you did this morning (available only between 5:00-7:59):

```bash
dawn "Reading"
```

Or use interactive mode:

```bash
dawn
# => What did you do?: Reading
```

### View streaks

```bash
dawn log
```

Output:

```
🔥 Current Streak:  12 days
🏆 Longest Streak:  34 days
📊 Streak Sum:      127 days
```

### View history

```bash
dawn history
```

Output:

```
2024-01-15 06:32  Reading
2024-01-14 05:58  Running
2024-01-13 06:15  English study
```

### Edit data

Open the JSON data file with your editor:

```bash
dawn edit
```

## Commands

| Command | Description |
|---------|-------------|
| `dawn [routine]` | Record a routine (5:00-7:59 only) |
| `dawn log` | Show streak statistics |
| `dawn history` | Show all records |
| `dawn edit` | Edit data file directly |
| `dawn help` | Show help |

## Constraints

- Recording is only available between **5:00 and 7:59**
- Only **one record per day** is allowed

## Data Storage

Data is stored in `~/.config/dawn/data.json`:

```json
{
  "records": [
    {
      "date": "2024-01-15",
      "time": "06:32",
      "routine": "Reading"
    }
  ]
}
```

## License

MIT
