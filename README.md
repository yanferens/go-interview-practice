# Go Interview Practice 🚀

Welcome to the **Go Interview Practice** repository! Master Go programming and ace your technical interviews with our interactive coding challenges.

---

## 🎨 **Visual Overview**

### 📋 Interactive Challenge Platform
Our comprehensive web interface provides everything you need to practice and master Go programming:

<div align="center">
  <img src="./images/challenges.png" alt="Go Interview Practice - Challenge Overview" width="90%">
  <p><em>📊 Complete challenge dashboard with difficulty levels, progress tracking, and performance metrics</em></p>
</div>

---

### 💻 **Code & Test Experience**

<div align="center">
  <img src="./images/challenge.png" alt="Go Interview Practice Web UI - challenge" width="48%" style="margin-right: 2%;">
  <img src="./images/result.png" alt="Go Interview Practice Web UI - result" width="48%">
</div>

<div align="center">
  <table>
    <tr>
      <td align="center" width="48%">
        <strong>🔧 Interactive Code Editor</strong><br>
        <em>Write, edit, and test your Go solutions<br>with syntax highlighting and real-time feedback</em>
      </td>
      <td width="4%"></td>
      <td align="center" width="48%">
        <strong>📈 Instant Results & Analytics</strong><br>
        <em>Get immediate test results, performance metrics,<br>and detailed execution analysis</em>
      </td>
    </tr>
  </table>
</div>

---

### 🏆 **Competitive Leaderboard**

<div align="center">
  <img src="./images/scoreboard.png" alt="Go Interview Practice - Main Leaderboard" width="90%">
  <p><em>🏅 Beautiful leaderboard showcasing top developers with challenge completion indicators, rankings, and achievements</em></p>
</div>

---

## 🏆 **Top 10 Leaderboard**

Our most accomplished Go developers, ranked by number of challenges completed:

> 📝 **Note**: The data below is automatically updated by GitHub Actions when challenge scoreboards change.


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
  <tbody>
    <tr style="background: white; border-bottom: 1px solid #dee2e6; transition: background-color 0.2s ease;">
      <td>
        <div class="rank-badge" style="background: linear-gradient(135deg, #ffd700, #ffed4e); color: #333; font-weight: bold;">1</div>
      </td>
      <td class="left-align">
        <div class="profile-cell">
          <img src="https://github.com/RezaSi.png" alt="RezaSi" class="profile-img">
          <div>
            <div class="username">RezaSi</div>
            <a href="https://github.com/RezaSi" class="github-link">
              <svg style="width: 12px; height: 12px; margin-right: 4px;" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/>
              </svg>View Profile
            </a>
          </div>
        </div>
      </td>
      <td>
        <div class="stat-number" style="color: #667eea;">12</div>
        <div class="stat-label">challenges</div>
      </td>
      <td class="rate-col">
        <div class="stat-number" style="color: #28a745; font-size: 16px;">42.9%</div>
        <div class="stat-label">complete</div>
      </td>
      <td class="achievement-col">
        <span class="achievement-badge" style="background: #6f42c1;">💪 Advanced</span>
      </td>
      <td class="progress-col left-align">
        <div class="progress-indicators">
          <span title="Challenge 1: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 2: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 3: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 4: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 5: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 6: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 7: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 8: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 9: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 10: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 11: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 12: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 13: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 14: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 15: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 16: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 17: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 18: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 19: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 20: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 21: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 22: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 23: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 24: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 25: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 26: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 27: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 28: Not completed" class="challenge-indicator challenge-not-completed">•</span>
        </div>
      </td>
    </tr>
    <tr style="background: #f8f9fa; border-bottom: 1px solid #dee2e6; transition: background-color 0.2s ease;">
      <td>
        <div class="rank-badge" style="background: linear-gradient(135deg, #c0c0c0, #e8e8e8); color: #333; font-weight: bold;">2</div>
      </td>
      <td class="left-align">
        <div class="profile-cell">
          <img src="https://github.com/AliNazariii.png" alt="AliNazariii" class="profile-img">
          <div>
            <div class="username">AliNazariii</div>
            <a href="https://github.com/AliNazariii" class="github-link">
              <svg style="width: 12px; height: 12px; margin-right: 4px;" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/>
              </svg>View Profile
            </a>
          </div>
        </div>
      </td>
      <td>
        <div class="stat-number" style="color: #667eea;">4</div>
        <div class="stat-label">challenges</div>
      </td>
      <td class="rate-col">
        <div class="stat-number" style="color: #28a745; font-size: 16px;">14.3%</div>
        <div class="stat-label">complete</div>
      </td>
      <td class="achievement-col">
        <span class="achievement-badge" style="background: #28a745;">🌱 Beginner</span>
      </td>
      <td class="progress-col left-align">
        <div class="progress-indicators">
          <span title="Challenge 1: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 2: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 3: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 4: Completed" class="challenge-indicator challenge-completed">✓</span><span title="Challenge 5: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 6: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 7: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 8: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 9: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 10: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 11: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 12: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 13: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 14: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 15: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 16: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 17: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 18: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 19: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 20: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 21: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 22: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 23: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 24: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 25: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 26: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 27: Not completed" class="challenge-indicator challenge-not-completed">•</span><span title="Challenge 28: Not completed" class="challenge-indicator challenge-not-completed">•</span>
        </div>
      </td>
    </tr>
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
</div>

*Updated automatically based on 28 available challenges*

### 🎯 **Challenge Progress Overview**

- **Total Challenges Available**: 28
- **Active Developers**: 2
- **Most Challenges Solved**: 12 by RezaSi

---

## 🌟 Key Features

- **Interactive Web UI** - Code, test, and submit solutions in your browser
- **Automated Testing** - Get immediate feedback on your solutions
- **Automated Scoreboards** - Solutions are automatically scored and ranked
- **Performance Analytics** - Track execution time and memory usage for your solutions
- **Comprehensive Learning** - Each challenge includes detailed explanations and resources
- **Progressive Difficulty** - From beginner to advanced Go concepts

## 🚀 Quick Start

### Option 1: Web UI (Recommended)

```bash
# Clone the repository
git clone https://github.com/yourusername/go-interview-practice.git
cd go-interview-practice

# Start the web interface
cd web-ui
go run main.go

# Open http://localhost:8080 in your browser
```

### Option 2: Command Line

```bash
# Set up a challenge workspace
./create_submission.sh 1  # For challenge #1

# Implement your solution in the editor of your choice

# Run tests
cd challenge-1
./run_tests.sh
```

## 📊 Scoreboards

Each challenge has its own scoreboard that tracks:
- Successful submissions by user
- Execution time rankings
- Code efficiency metrics
- Completion dates

View global and per-challenge scoreboards in the Web UI to compare your solutions with others.

## 📚 Challenge Categories

### 🟢 Beginner
Perfect for those new to Go or brushing up on fundamentals
- **[Challenge 1](./challenge-1)**: Sum of Two Numbers
- **[Challenge 2](./challenge-2)**: Reverse a String
- **[Challenge 3](./challenge-3)**: Employee Data Management
- **[Challenge 6](./challenge-6)**: Word Frequency Counter
- **[Challenge 18](./challenge-18)**: Temperature Converter
- **[Challenge 21](./challenge-21)**: Binary Search Implementation
- **[Challenge 22](./challenge-22)**: Greedy Coin Change

### 🟠 Intermediate
For developers familiar with Go who want to deepen their knowledge
- **[Challenge 4](./challenge-4)**: Concurrent Graph BFS Queries
- **[Challenge 5](./challenge-5)**: HTTP Authentication Middleware
- **[Challenge 7](./challenge-7)**: Bank Account with Error Handling
- **[Challenge 10](./challenge-10)**: Polymorphic Shape Calculator
- **[Challenge 13](./challenge-13)**: SQL Database Operations
- **[Challenge 14](./challenge-14)**: Microservices with gRPC
- **[Challenge 16](./challenge-16)**: Performance Optimization
- **[Challenge 17](./challenge-17)**: Interactive Debugging Tutorial
- **[Challenge 19](./challenge-19)**: Slice Operations
- **[Challenge 23](./challenge-23)**: String Pattern Matching
- **[Challenge 27](./challenge-27)**: Go Generics Data Structures

### 🔴 Advanced
Challenging problems that test mastery of Go and computer science concepts
- **[Challenge 8](./challenge-8)**: Chat Server with Channels
- **[Challenge 9](./challenge-9)**: RESTful Book Management API
- **[Challenge 11](./challenge-11)**: Concurrent Web Content Aggregator
- **[Challenge 12](./challenge-12)**: File Processing Pipeline
- **[Challenge 15](./challenge-15)**: OAuth2 Authentication
- **[Challenge 20](./challenge-20)**: Circuit Breaker Pattern
- **[Challenge 24](./challenge-24)**: Dynamic Programming - Longest Increasing Subsequence
- **[Challenge 25](./challenge-25)**: Graph Algorithms - Shortest Path
- **[Challenge 26](./challenge-26)**: Regular Expression Text Processor
- **[Challenge 28](./challenge-28)**: Cache Implementation with Multiple Eviction Policies

## 💡 How to Use This Repository

### 1. Explore Challenges
Browse challenges through the web UI or in the code repository. Each challenge includes:
- Detailed problem statement
- Function signature to implement
- Comprehensive test cases
- Learning resources

### 2. Implement Your Solution
Write code that solves the challenge requirements and passes all test cases.

### 3. Test & Refine
Use the built-in testing tools to validate your solution, then refine it for:
- Correctness
- Efficiency
- Code quality

### 4. Submit & Compare
Submit your passing solution to be added to the scoreboard:
- Your solution is automatically tested and scored
- Execution time and resource usage are recorded
- Your solution is ranked among other submissions
- Access detailed performance metrics to optimize further

### 5. Learn & Progress
Review the learning materials to deepen your understanding of the concepts used.

## 🤝 Contributing

We welcome contributions! To add a new challenge:

1. Fork the repository
2. Create a new challenge following our template structure
3. Submit a pull request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Happy Coding!** 💻
