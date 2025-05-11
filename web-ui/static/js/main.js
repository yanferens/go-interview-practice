// Common JavaScript utilities for Go Interview Practice web UI

// Helper for formatting timestamps
function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
}

// Helper for truncating text
function truncateText(text, maxLength = 150) {
    if (text.length <= maxLength) return text;
    return text.substr(0, maxLength) + '...';
}

// Helper for formatting execution time
function formatExecutionTime(ms) {
    if (ms < 1000) {
        return ms + 'ms';
    } else {
        return (ms / 1000).toFixed(2) + 's';
    }
}

// Initialize syntax highlighting for code blocks
function initSyntaxHighlighting() {
    document.querySelectorAll('pre code').forEach((el) => {
        hljs.highlightElement(el);
    });
}

// Initialize Markdown parsing
function renderMarkdown(markdownText, targetElement) {
    if (!markdownText || !targetElement) return;
    
    // Set options for the markdown parser
    marked.setOptions({
        gfm: true,      // GitHub flavored markdown
        breaks: true,   // Line breaks are rendered
        highlight: function(code, lang) {
            const language = hljs.getLanguage(lang) ? lang : 'plaintext';
            return hljs.highlight(code, { language }).value;
        }
    });
    
    // Parse and render markdown
    targetElement.innerHTML = marked.parse(markdownText);
    
    // Initialize syntax highlighting on code blocks
    initSyntaxHighlighting();
}

// Helper for creating a code editor
function createEditor(elementId, code, isReadOnly = false) {
    const editor = ace.edit(elementId);
    editor.setTheme("ace/theme/chrome");
    editor.session.setMode("ace/mode/golang");
    editor.setValue(code || '');
    editor.setReadOnly(isReadOnly);
    editor.clearSelection();
    editor.setOptions({
        fontSize: "14px",
        showPrintMargin: false,
        showGutter: true,
        highlightActiveLine: true,
        wrap: true
    });
    return editor;
}

// Save editor content in localStorage to prevent loss on page refresh
function saveEditorContent(key, content) {
    localStorage.setItem(`editor_${key}`, content);
}

// Load editor content from localStorage if available
function loadEditorContent(key) {
    return localStorage.getItem(`editor_${key}`);
}

// Format the test output for display
function formatTestOutput(output) {
    if (!output) return '';
    
    // Colorize PASS/FAIL in the output
    return output
        .replace(/PASS/g, '<span class="text-success">PASS</span>')
        .replace(/FAIL/g, '<span class="text-danger">FAIL</span>')
        .replace(/--- FAIL/g, '<span class="text-danger">--- FAIL</span>')
        .replace(/--- PASS/g, '<span class="text-success">--- PASS</span>');
}

// Handle form submissions with AJAX
function handleFormSubmit(formElement, successCallback, errorCallback) {
    formElement.addEventListener('submit', function(e) {
        e.preventDefault();
        const formData = new FormData(formElement);
        const jsonData = {};
        
        for (const [key, value] of formData.entries()) {
            jsonData[key] = value;
        }
        
        fetch(formElement.action, {
            method: formElement.method || 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData)
        })
        .then(response => response.json())
        .then(data => {
            if (typeof successCallback === 'function') {
                successCallback(data);
            }
        })
        .catch(error => {
            if (typeof errorCallback === 'function') {
                errorCallback(error);
            } else {
                console.error('Error:', error);
            }
        });
    });
}

// Custom template function for truncating text in templates
function truncateDescription(text, maxLength = 100) {
    if (!text) return '';
    const plainText = text.replace(/<[^>]*>/g, '');
    if (plainText.length <= maxLength) return plainText;
    return plainText.substr(0, maxLength) + '...';
}

// Initialize common page elements
document.addEventListener('DOMContentLoaded', function() {
    // Initialize tooltips
    const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl);
    });
    
    // Initialize syntax highlighting
    initSyntaxHighlighting();
    
    // Handle username persistence
    const usernameInput = document.getElementById('username');
    if (usernameInput) {
        // Load saved username
        const savedUsername = localStorage.getItem('githubUsername');
        if (savedUsername) {
            usernameInput.value = savedUsername;
        }
        
        // Save username when changed
        usernameInput.addEventListener('change', function() {
            localStorage.setItem('githubUsername', this.value);
        });
    }
}); 