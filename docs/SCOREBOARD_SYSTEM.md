# 🏆 Scoreboard System Documentation

This document explains how the Go Interview Practice scoreboard system works, including both individual challenge scoreboards and the main leaderboard.

## 📊 Overview

The scoreboard system consists of two levels:

1. **Individual Challenge Scoreboards** - Track submissions for each specific challenge
2. **Main Leaderboard** - Aggregates data across all challenges to show top performers

## 🔄 How It Works

### Individual Challenge Scoreboards

Each challenge directory contains a `SCOREBOARD.md` file that tracks successful submissions:

```
challenge-1/
├── SCOREBOARD.md
├── README.md
└── ...
```

**Format:**
```markdown
# Scoreboard for challenge-1
| Username   | Passed Tests | Total Tests |
|------------|--------------|-------------|
| RezaSi     | 6            | 6           |
| AliNazariii| 6            | 6           |
```

### Main Leaderboard (README.md)

The main leaderboard aggregates completion data from all challenges and displays:
- Top 10 developers by number of challenges solved
- Completion rates and achievement badges
- Overall statistics

## 🤖 Automated Updates

### GitHub Actions Workflow

The system uses two GitHub Actions workflows:

#### 1. Update Challenge Scoreboards (`.github/workflows/update-scoreboards.yml`)
- **Triggers**: On push to main branch
- **Process**: 
  - Runs tests for all submissions in each challenge
  - Updates individual `SCOREBOARD.md` files
  - Calls main scoreboard update
  - Commits and pushes changes

#### 2. Update Main Scoreboard (`.github/workflows/update-main-scoreboard.yml`)
- **Triggers**: 
  - When challenge scoreboards change
  - Daily at 00:00 UTC (scheduled)
  - Manual dispatch
- **Process**:
  - Aggregates data from all challenge scoreboards
  - Updates the main leaderboard in README.md
  - Commits and pushes changes

### Automatic Triggering

The main scoreboard updates automatically when:
- Any `challenge-*/SCOREBOARD.md` file changes
- The daily scheduled workflow runs
- Manual workflow dispatch is triggered

## 📈 Achievement System

Developers earn achievement badges based on completion count:

| Badge | Name | Requirements |
|-------|------|-------------|
| 🔥 | **Master** | 20+ challenges completed |
| ⭐ | **Expert** | 15+ challenges completed |
| 💪 | **Advanced** | 10+ challenges completed |
| 🚀 | **Intermediate** | 5+ challenges completed |
| 🌱 | **Beginner** | 1+ challenges completed |

## 🛠 Manual Operations

### Update Main Scoreboard Manually

Run the provided script:
```bash
./scripts/update_scoreboard.sh
```

This will:
1. Analyze all challenge scoreboards
2. Aggregate completion data
3. Update the main leaderboard in README.md
4. Display summary statistics

### Update Specific Challenge Scoreboard

Navigate to a challenge directory and run tests:
```bash
cd challenge-1
./run_tests.sh username
```

## 📁 File Structure

```
.
├── README.md                           # Contains main leaderboard
├── scripts/
│   ├── generate_main_scoreboard.py     # Python script to generate leaderboard
│   └── update_scoreboard.sh            # Shell script for manual updates
├── .github/workflows/
│   ├── update-scoreboards.yml          # Update individual scoreboards
│   └── update-main-scoreboard.yml      # Update main leaderboard
└── challenge-*/
    └── SCOREBOARD.md                   # Individual challenge scoreboards
```

## 🔧 Technical Details

### Data Aggregation Logic

The `generate_main_scoreboard.py` script:

1. **Scans** all `challenge-*/SCOREBOARD.md` files
2. **Parses** markdown tables to extract usernames
3. **Counts** unique challenge completions per user
4. **Sorts** users by completion count (descending) then by username
5. **Generates** markdown table with rankings and statistics
6. **Updates** README.md with new leaderboard

### Error Handling

- Handles missing or malformed scoreboard files gracefully
- Skips invalid usernames (empty, numeric, placeholder values)
- Provides detailed logging for debugging
- Fails gracefully without breaking the repository

## 🎯 Future Enhancements

Potential improvements to consider:

- **Performance Metrics**: Include execution time rankings
- **Difficulty Weighting**: Weight challenges by difficulty level
- **Historical Tracking**: Track progress over time
- **Team Scoreboards**: Support for team-based challenges
- **Detailed Statistics**: Per-user challenge completion details

## 🤝 Contributing

To modify the scoreboard system:

1. **Challenge Scoreboards**: Update format in individual `run_tests.sh` scripts
2. **Main Leaderboard**: Modify `scripts/generate_main_scoreboard.py`
3. **Workflows**: Update `.github/workflows/` files for automation changes
4. **Documentation**: Update this file and README.md accordingly

For questions or suggestions about the scoreboard system, please open an issue or contribute improvements! 