// DOM Elements
const themeToggle = document.getElementById('themeToggle');
const inputText = document.getElementById('inputText');
const outputText = document.getElementById('outputText');
const encodeBtn = document.getElementById('encodeBtn');
const decodeBtn = document.getElementById('decodeBtn');
const clearInput = document.getElementById('clearInput');
const clearOutput = document.getElementById('clearOutput');
const copyInput = document.getElementById('copyInput');
const copyOutput = document.getElementById('copyOutput');
const toast = document.getElementById('toast');

// Theme Management
function initTheme() {
    const savedTheme = localStorage.getItem('theme') || 'dark';
    document.documentElement.setAttribute('data-theme', savedTheme);
}

themeToggle.addEventListener('click', () => {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    document.documentElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
});

// Toast Notification
function showToast(message, type = 'success') {
    toast.textContent = message;
    toast.className = `toast ${type} show`;
    
    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}

// Brainfuck Encoder
function textToBrainfuck(text) {
    if (!text) {
        showToast('Silakan masukkan teks terlebih dahulu!', 'error');
        return '';
    }
    
    let brainfuck = '';
    let currentValue = 0;
    
    for (let i = 0; i < text.length; i++) {
        const charCode = text.charCodeAt(i);
        const diff = charCode - currentValue;
        
        if (diff > 0) {
            // Optimized increment
            if (diff > 10) {
                const sqrt = Math.floor(Math.sqrt(diff));
                const remainder = diff - (sqrt * sqrt);
                brainfuck += '+'.repeat(sqrt) + '[>' + '+'.repeat(sqrt) + '<-]>' + '+'.repeat(remainder);
                currentValue = charCode;
            } else {
                brainfuck += '+'.repeat(diff);
                currentValue = charCode;
            }
        } else if (diff < 0) {
            brainfuck += '-'.repeat(Math.abs(diff));
            currentValue = charCode;
        }
        
        brainfuck += '.';
    }
    
    return brainfuck;
}

// Brainfuck Decoder/Interpreter
function executeBrainfuck(code) {
    if (!code) {
        showToast('Silakan masukkan kode Brainfuck terlebih dahulu!', 'error');
        return '';
    }
    
    // Remove non-brainfuck characters
    const cleanCode = code.replace(/[^><+\-.,\[\]]/g, '');
    
    const memory = new Array(30000).fill(0);
    let pointer = 0;
    let codePointer = 0;
    let output = '';
    
    const loopStack = [];
    const loopMap = {};
    
    // Build loop map for jumping
    for (let i = 0; i < cleanCode.length; i++) {
        if (cleanCode[i] === '[') {
            loopStack.push(i);
        } else if (cleanCode[i] === ']') {
            if (loopStack.length === 0) {
                showToast('Error: Unmatched ]', 'error');
                return '';
            }
            const start = loopStack.pop();
            loopMap[start] = i;
            loopMap[i] = start;
        }
    }
    
    if (loopStack.length > 0) {
        showToast('Error: Unmatched [', 'error');
        return '';
    }
    
    // Execute brainfuck code
    let iterations = 0;
    const maxIterations = 1000000; // Prevent infinite loops
    
    while (codePointer < cleanCode.length) {
        if (iterations++ > maxIterations) {
            showToast('Error: Execution timeout (infinite loop?)', 'error');
            return output + '\n[TIMEOUT]';
        }
        
        const command = cleanCode[codePointer];
        
        switch (command) {
            case '>':
                pointer++;
                if (pointer >= memory.length) {
                    showToast('Error: Memory overflow', 'error');
                    return output;
                }
                break;
                
            case '<':
                pointer--;
                if (pointer < 0) {
                    showToast('Error: Memory underflow', 'error');
                    return output;
                }
                break;
                
            case '+':
                memory[pointer] = (memory[pointer] + 1) % 256;
                break;
                
            case '-':
                memory[pointer] = (memory[pointer] - 1 + 256) % 256;
                break;
                
            case '.':
                output += String.fromCharCode(memory[pointer]);
                break;
                
            case ',':
                // Input not supported in this implementation
                memory[pointer] = 0;
                break;
                
            case '[':
                if (memory[pointer] === 0) {
                    codePointer = loopMap[codePointer];
                }
                break;
                
            case ']':
                if (memory[pointer] !== 0) {
                    codePointer = loopMap[codePointer];
                }
                break;
        }
        
        codePointer++;
    }
    
    return output;
}

// Event Handlers
encodeBtn.addEventListener('click', () => {
    const input = inputText.value;
    const result = textToBrainfuck(input);
    
    if (result) {
        outputText.value = result;
        showToast('Teks berhasil di-encode ke Brainfuck!', 'success');
    }
});

decodeBtn.addEventListener('click', () => {
    const input = inputText.value;
    const result = executeBrainfuck(input);
    
    if (result) {
        outputText.value = result;
        showToast('Kode Brainfuck berhasil dijalankan!', 'success');
    }
});

// Clear buttons
clearInput.addEventListener('click', () => {
    inputText.value = '';
    inputText.focus();
    showToast('Input dibersihkan', 'success');
});

clearOutput.addEventListener('click', () => {
    outputText.value = '';
    showToast('Output dibersihkan', 'success');
});

// Copy buttons
copyInput.addEventListener('click', async () => {
    if (!inputText.value) {
        showToast('Tidak ada yang dapat disalin!', 'error');
        return;
    }
    
    try {
        await navigator.clipboard.writeText(inputText.value);
        showToast('Input disalin ke clipboard!', 'success');
    } catch (err) {
        showToast('Gagal menyalin ke clipboard', 'error');
    }
});

copyOutput.addEventListener('click', async () => {
    if (!outputText.value) {
        showToast('Tidak ada yang dapat disalin!', 'error');
        return;
    }
    
    try {
        await navigator.clipboard.writeText(outputText.value);
        showToast('Output disalin ke clipboard!', 'success');
    } catch (err) {
        showToast('Gagal menyalin ke clipboard', 'error');
    }
});

// Keyboard shortcuts
inputText.addEventListener('keydown', (e) => {
    // Ctrl/Cmd + Enter to encode
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        e.preventDefault();
        encodeBtn.click();
    }
});

// Initialize
initTheme();

// Demo on load (optional)
window.addEventListener('load', () => {
    // You can add a demo text here if you want
    // inputText.value = 'Hello World!';
});
