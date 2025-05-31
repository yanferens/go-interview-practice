#!/usr/bin/env python3
"""
Script to generate the main scoreboard for README.md by aggregating data from all challenge scoreboards.
"""

import os
import re
import sys
from collections import defaultdict
from pathlib import Path


def parse_scoreboard_file(filepath):
    """Parse a SCOREBOARD.md file and extract usernames who completed the challenge."""
    users = set()
    
    try:
        with open(filepath, 'r') as f:
            content = f.read()
        
        # Split into lines and process
        lines = content.strip().split('\n')
        
        for line in lines:
            # Skip header and separator lines
            if not line.strip() or 'Username' in line or '---' in line or line.startswith('#'):
                continue
            
            # Parse table row
            if '|' in line:
                parts = [part.strip() for part in line.split('|')]
                if len(parts) >= 3:
                    username = parts[1]
                    
                    # Skip empty usernames or placeholders
                    if username and username != '------' and not username.isdigit():
                        users.add(username)
    
    except FileNotFoundError:
        pass
    except Exception as e:
        print(f"Error parsing {filepath}: {e}", file=sys.stderr)
    
    return users


def get_challenge_title(challenge_dir):
    """Extract challenge title from README.md or directory name."""
    readme_path = os.path.join(challenge_dir, 'README.md')
    
    try:
        with open(readme_path, 'r') as f:
            content = f.read()
        
        # Look for title in various formats
        title_patterns = [
            r'^#\s+(.+?)$',  # # Title
            r'^\*\*(.+?)\*\*',  # **Title**
            r'Challenge \d+:\s*(.+?)$'  # Challenge N: Title
        ]
        
        lines = content.split('\n')
        for line in lines:
            line = line.strip()
            for pattern in title_patterns:
                match = re.search(pattern, line, re.MULTILINE)
                if match:
                    title = match.group(1).strip()
                    # Clean up title
                    title = re.sub(r'Challenge \d+:\s*', '', title)
                    return title
    
    except FileNotFoundError:
        pass
    
    # Fallback to directory name
    return challenge_dir.replace('challenge-', 'Challenge ')


def generate_main_scoreboard():
    """Generate the main scoreboard by aggregating all challenge scoreboards."""
    
    # Dictionary to track user completion counts
    user_completions = defaultdict(lambda: {'count': 0, 'challenges': []})
    
    # Find all challenge directories
    current_dir = Path('.')
    challenge_dirs = sorted([d for d in current_dir.iterdir() 
                           if d.is_dir() and d.name.startswith('challenge-')])
    
    print(f"Found {len(challenge_dirs)} challenge directories")
    
    # Process each challenge
    for challenge_dir in challenge_dirs:
        challenge_num = challenge_dir.name.replace('challenge-', '')
        scoreboard_path = challenge_dir / 'SCOREBOARD.md'
        
        if scoreboard_path.exists():
            users = parse_scoreboard_file(scoreboard_path)
            challenge_title = get_challenge_title(str(challenge_dir))
            
            print(f"Challenge {challenge_num}: {len(users)} users completed")
            
            for user in users:
                user_completions[user]['count'] += 1
                user_completions[user]['challenges'].append({
                    'num': int(challenge_num),
                    'title': challenge_title
                })
    
    # Sort users by completion count (descending) and then by username
    sorted_users = sorted(user_completions.items(), 
                         key=lambda x: (-x[1]['count'], x[0]))
    
    # Generate the markdown table
    markdown_lines = [
        "## 🏆 **Top 10 Leaderboard**",
        "",
        "Our most accomplished Go developers, ranked by number of challenges completed:",
        "",
        "| Rank | 👤 Developer | 🎯 Challenges Solved | 📊 Completion Rate | 🏅 Achievement |",
        "|------|-------------|---------------------|-------------------|----------------|"
    ]
    
    total_challenges = len(challenge_dirs)
    
    for i, (username, data) in enumerate(sorted_users[:10], 1):
        count = data['count']
        completion_rate = f"{(count / total_challenges * 100):.1f}%"
        
        # Determine achievement badge
        if count >= 20:
            achievement = "🔥 **Master**"
        elif count >= 15:
            achievement = "⭐ **Expert**"
        elif count >= 10:
            achievement = "💪 **Advanced**"
        elif count >= 5:
            achievement = "🚀 **Intermediate**"
        else:
            achievement = "🌱 **Beginner**"
        
        markdown_lines.append(
            f"| **{i}** | **{username}** | **{count}** / {total_challenges} | {completion_rate} | {achievement} |"
        )
    
    # Add summary information
    markdown_lines.extend([
        "",
        f"*Updated automatically based on {total_challenges} available challenges*",
        "",
        "### 🎯 **Challenge Progress Overview**",
        "",
        f"- **Total Challenges Available**: {total_challenges}",
        f"- **Active Developers**: {len(user_completions)}",
        f"- **Most Challenges Solved**: {sorted_users[0][1]['count'] if sorted_users else 0} by {sorted_users[0][0] if sorted_users else 'N/A'}",
        "",
        "---",
        ""
    ])
    
    return '\n'.join(markdown_lines)


def update_readme_with_scoreboard(scoreboard_content):
    """Update README.md with the new scoreboard content."""
    
    readme_path = 'README.md'
    
    try:
        with open(readme_path, 'r') as f:
            content = f.read()
    except FileNotFoundError:
        print("README.md not found!", file=sys.stderr)
        return False
    
    # Define markers for the scoreboard section
    start_marker = "## 🏆 **Top 10 Leaderboard**"
    end_marker = "## 🌟 Key Features"
    
    # Find the positions of markers
    start_pos = content.find(start_marker)
    end_pos = content.find(end_marker)
    
    if start_pos == -1 or end_pos == -1:
        # If markers don't exist, insert before Key Features section
        key_features_pos = content.find("## 🌟 Key Features")
        if key_features_pos == -1:
            print("Could not find insertion point in README.md", file=sys.stderr)
            return False
        
        # Insert the new scoreboard before Key Features
        new_content = (content[:key_features_pos] + 
                      scoreboard_content + '\n' + 
                      content[key_features_pos:])
    else:
        # Replace existing scoreboard section
        new_content = (content[:start_pos] + 
                      scoreboard_content + '\n' + 
                      content[end_pos:])
    
    # Write the updated content
    try:
        with open(readme_path, 'w') as f:
            f.write(new_content)
        print("README.md updated successfully!")
        return True
    except Exception as e:
        print(f"Error writing to README.md: {e}", file=sys.stderr)
        return False


def main():
    """Main function to generate and update the scoreboard."""
    print("Generating main scoreboard...")
    
    scoreboard_content = generate_main_scoreboard()
    
    if update_readme_with_scoreboard(scoreboard_content):
        print("Main scoreboard updated successfully!")
        return 0
    else:
        print("Failed to update main scoreboard!", file=sys.stderr)
        return 1


if __name__ == "__main__":
    sys.exit(main()) 