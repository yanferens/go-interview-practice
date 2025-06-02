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
    """Parse a SCOREBOARD.md file and extract usernames who completed the challenge (passed ALL tests)."""
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
                if len(parts) >= 4:  # Username | Passed Tests | Total Tests | (optional extra columns)
                    username = parts[1]
                    passed_tests_str = parts[2]
                    total_tests_str = parts[3]
                    
                    # Skip empty usernames or placeholders
                    if not username or username == '------' or username.isdigit():
                        continue
                    
                    try:
                        # Extract numbers from test counts (handles formats like "6", "6 tests", etc.)
                        passed_tests = int(''.join(filter(str.isdigit, passed_tests_str)))
                        total_tests = int(''.join(filter(str.isdigit, total_tests_str)))
                        
                        # Only count as completed if ALL tests passed
                        if passed_tests > 0 and passed_tests == total_tests:
                            users.add(username)
                            print(f"  ✅ {username}: {passed_tests}/{total_tests} tests passed (COMPLETED)")
                        else:
                            print(f"  ❌ {username}: {passed_tests}/{total_tests} tests passed (incomplete)")
                            
                    except (ValueError, TypeError):
                        # If we can't parse test numbers, skip this entry
                        print(f"  ⚠️  {username}: Could not parse test results")
                        continue
    
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
    
    # Generate the HTML leaderboard
    markdown_lines = [
        "## 🏆 **Top 10 Leaderboard**",
        "",
        "Our most accomplished Go developers, ranked by number of challenges completed:",
        "",
        "> 📝 **Note**: The data below is automatically updated by GitHub Actions when challenge scoreboards change.",
        "",
    ]
    
    total_challenges = len(challenge_dirs)
    
    if sorted_users:
        # Generate HTML table with styling
        html_table = generate_html_leaderboard(sorted_users[:10], total_challenges, challenge_dirs)
        markdown_lines.append(html_table)
    else:
        markdown_lines.extend([
            "No completed challenges yet. Be the first to solve a challenge!",
            ""
        ])
    
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


def generate_html_leaderboard(top_users, total_challenges, challenge_dirs):
    """Generate a beautiful HTML leaderboard table."""
    
    # Get list of challenge numbers for indicators
    challenge_numbers = sorted([int(d.name.replace('challenge-', '')) for d in challenge_dirs])
    
    html = """
<style>
  .leaderboard-table {
    border-collapse: collapse; 
    width: 100%; 
    max-width: 1000px; 
    margin: 20px 0; 
    box-shadow: 0 4px 15px rgba(102, 126, 234, 0.1); 
    border-radius: 12px; 
    overflow: hidden;
  }
  .leaderboard-table th, .leaderboard-table td {
    border: none;
  }
  .leaderboard-table thead tr {
    background: linear-gradient(135deg, #667eea, #764ba2); 
    color: white;
  }
  .leaderboard-table th {
    padding: 15px 10px; 
    text-align: center; 
    font-weight: 600;
  }
  .leaderboard-table th.left-align {
    text-align: left; 
    padding: 15px 15px;
  }
  .leaderboard-table td {
    padding: 15px 10px; 
    text-align: center;
  }
  .leaderboard-table td.left-align {
    padding: 15px; 
    text-align: left;
  }
  .rank-badge {
    width: 35px; 
    height: 35px; 
    border-radius: 50%; 
    display: flex; 
    align-items: center; 
    justify-content: center; 
    margin: 0 auto; 
    font-size: 14px;
  }
  .profile-cell {
    display: flex; 
    align-items: center;
  }
  .profile-img {
    width: 40px; 
    height: 40px; 
    border-radius: 50%; 
    border: 2px solid #fff; 
    box-shadow: 0 2px 8px rgba(0,0,0,0.1); 
    margin-right: 12px;
  }
  .username {
    font-weight: 600; 
    color: #2c3e50; 
    font-size: 16px;
  }
  .github-link {
    color: #6c757d; 
    text-decoration: none; 
    font-size: 12px;
  }
  .stat-number {
    font-weight: 700; 
    font-size: 20px;
  }
  .stat-label {
    color: #6c757d; 
    font-size: 12px;
  }
  .achievement-badge {
    color: white; 
    padding: 6px 12px; 
    border-radius: 20px; 
    font-size: 12px; 
    font-weight: 500;
  }
  .progress-indicators {
    line-height: 1.2; 
    max-width: 300px;
  }
  .challenge-indicator {
    display: inline-block; 
    width: 18px; 
    height: 18px; 
    border-radius: 3px; 
    margin: 1px; 
    text-align: center; 
    line-height: 18px; 
    font-size: 10px;
  }
  .challenge-completed {
    background: #28a745; 
    color: white; 
    font-weight: bold;
  }
  .challenge-not-completed {
    background: #e9ecef; 
    color: #6c757d;
  }
  
  /* Responsive breakpoints */
  @media (max-width: 968px) {
    .progress-col {
      display: none;
    }
    .leaderboard-table {
      max-width: 700px;
    }
  }
  
  @media (max-width: 768px) {
    .achievement-col {
      display: none;
    }
    .leaderboard-table {
      max-width: 500px;
    }
    .leaderboard-table th, .leaderboard-table td {
      padding: 12px 8px;
    }
    .leaderboard-table th.left-align, .leaderboard-table td.left-align {
      padding: 12px;
    }
    .profile-img {
      width: 32px; 
      height: 32px;
    }
    .username {
      font-size: 14px;
    }
    .stat-number {
      font-size: 18px;
    }
  }
  
  @media (max-width: 480px) {
    .rate-col {
      display: none;
    }
    .leaderboard-table {
      max-width: 100%;
    }
    .leaderboard-table th, .leaderboard-table td {
      padding: 10px 6px;
    }
    .leaderboard-table th.left-align, .leaderboard-table td.left-align {
      padding: 10px 8px;
    }
    .rank-badge {
      width: 28px; 
      height: 28px; 
      font-size: 12px;
    }
    .profile-img {
      width: 28px; 
      height: 28px; 
      margin-right: 8px;
    }
    .username {
      font-size: 13px;
    }
    .github-link {
      font-size: 10px;
    }
    .stat-number {
      font-size: 16px;
    }
    .stat-label {
      font-size: 10px;
    }
  }
</style>

<div align="center">
<table class="leaderboard-table">
  <thead>
    <tr>
      <th>🏅 Rank</th>
      <th class="left-align">👤 Developer</th>
      <th>🎯 Solved</th>
      <th class="rate-col">📊 Rate</th>
      <th class="achievement-col">🏅 Achievement</th>
      <th class="progress-col left-align">📈 Progress</th>
    </tr>
  </thead>
  <tbody>"""

    for i, (username, data) in enumerate(top_users, 1):
        count = data['count']
        completion_rate = f"{(count / total_challenges * 100):.1f}%"
        
        # Determine achievement badge
        if count >= 20:
            achievement = "🔥 Master"
            achievement_color = "#dc3545"
        elif count >= 15:
            achievement = "⭐ Expert"
            achievement_color = "#fd7e14"
        elif count >= 10:
            achievement = "💪 Advanced"
            achievement_color = "#6f42c1"
        elif count >= 5:
            achievement = "🚀 Intermediate"
            achievement_color = "#20c997"
        else:
            achievement = "🌱 Beginner"
            achievement_color = "#28a745"
        
        # Rank badge styling
        if i == 1:
            rank_style = "background: linear-gradient(135deg, #ffd700, #ffed4e); color: #333; font-weight: bold;"
        elif i <= 3:
            rank_style = "background: linear-gradient(135deg, #c0c0c0, #e8e8e8); color: #333; font-weight: bold;"
        elif i <= 10:
            rank_style = "background: linear-gradient(135deg, #cd7f32, #daa520); color: white; font-weight: bold;"
        else:
            rank_style = "background: linear-gradient(135deg, #6c757d, #495057); color: white; font-weight: bold;"
        
        # Generate challenge indicators
        completed_challenges = {ch['num'] for ch in data['challenges']}
        indicators = ""
        for ch_num in challenge_numbers:
            if ch_num in completed_challenges:
                indicators += f'<span title="Challenge {ch_num}: Completed" class="challenge-indicator challenge-completed">✓</span>'
            else:
                indicators += f'<span title="Challenge {ch_num}: Not completed" class="challenge-indicator challenge-not-completed">•</span>'
        
        # Row styling with alternating colors
        row_bg = "#f8f9fa" if i % 2 == 0 else "white"
        
        html += f"""
    <tr style="background: {row_bg}; border-bottom: 1px solid #dee2e6; transition: background-color 0.2s ease;">
      <td>
        <div class="rank-badge" style="{rank_style}">{i}</div>
      </td>
      <td class="left-align">
        <div class="profile-cell">
          <img src="https://github.com/{username}.png" alt="{username}" class="profile-img">
          <div>
            <div class="username">{username}</div>
            <a href="https://github.com/{username}" class="github-link">
              <svg style="width: 12px; height: 12px; margin-right: 4px;" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/>
              </svg>View Profile
            </a>
          </div>
        </div>
      </td>
      <td>
        <div class="stat-number" style="color: #667eea;">{count}</div>
        <div class="stat-label">challenges</div>
      </td>
      <td class="rate-col">
        <div class="stat-number" style="color: #28a745; font-size: 16px;">{completion_rate}</div>
        <div class="stat-label">complete</div>
      </td>
      <td class="achievement-col">
        <span class="achievement-badge" style="background: {achievement_color};">{achievement}</span>
      </td>
      <td class="progress-col left-align">
        <div class="progress-indicators">
          {indicators}
        </div>
      </td>
    </tr>"""

    html += """
  </tbody>
</table>
</div>

<div align="center" style="margin-top: 20px;">
  <div style="display: inline-flex; align-items: center; background: #f8f9fa; padding: 12px 20px; border-radius: 25px; border: 1px solid #dee2e6; flex-wrap: wrap; justify-content: center;">
    <span style="color: #6c757d; font-size: 14px; margin-right: 15px; margin-bottom: 5px;"><strong>Legend:</strong></span>
    <span class="challenge-indicator challenge-completed" style="margin-right: 8px; margin-bottom: 5px;">✓</span>
    <span style="color: #6c757d; font-size: 12px; margin-right: 15px; margin-bottom: 5px;">Completed</span>
    <span class="challenge-indicator challenge-not-completed" style="margin-right: 8px; margin-bottom: 5px;">•</span>
    <span style="color: #6c757d; font-size: 12px; margin-bottom: 5px;">Not completed</span>
  </div>
</div>"""

    return html


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