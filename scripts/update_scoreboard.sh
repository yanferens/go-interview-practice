#!/bin/bash

# Script to manually update the main scoreboard in README.md
# This aggregates data from all challenge scoreboards

set -e

echo "🏆 Go Interview Practice - Main Scoreboard Updater"
echo "=================================================="
echo ""

# Check if we're in the right directory
if [ ! -f "README.md" ]; then
    echo "❌ Error: README.md not found. Please run this script from the repository root."
    exit 1
fi

# Check if Python 3 is available
if ! command -v python3 &> /dev/null; then
    echo "❌ Error: Python 3 is required but not installed."
    exit 1
fi

# Check if the script directory exists
if [ ! -d "scripts" ]; then
    echo "❌ Error: scripts directory not found."
    exit 1
fi

# Check if the generator script exists
if [ ! -f "scripts/generate_main_scoreboard.py" ]; then
    echo "❌ Error: Main scoreboard generator script not found."
    exit 1
fi

echo "📊 Analyzing challenge scoreboards..."
echo ""

# Count challenges with submissions
challenges_with_submissions=$(find . -name "SCOREBOARD.md" -exec grep -l -v "^#\|^|\s*Username\|^|\s*---" {} \; 2>/dev/null | wc -l | xargs)
total_challenges=$(find . -name "SCOREBOARD.md" | wc -l | xargs)

echo "📈 Found:"
echo "  - Total challenges: $total_challenges"
echo "  - Challenges with submissions: $challenges_with_submissions"
echo ""

# Run the main scoreboard generator
echo "🔄 Generating main scoreboard..."
python3 scripts/generate_main_scoreboard.py

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Main scoreboard updated successfully!"
    echo ""
    echo "📝 The leaderboard in README.md has been updated with:"
    echo "  - Latest completion statistics"
    echo "  - Updated achievement badges"
    echo "  - Current completion rates"
    echo "  - Top 10 developer rankings"
    echo ""
    echo "🎯 Next steps:"
    echo "  1. Review the updated README.md"
    echo "  2. Commit and push changes if satisfied"
    echo "  3. The scoreboard will auto-update on future submissions"
    echo ""
else
    echo ""
    echo "❌ Failed to update main scoreboard!"
    echo "Please check the error messages above and try again."
    exit 1
fi 