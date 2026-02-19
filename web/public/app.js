// Kamadhenu Web App
let fortunes = [];
let cows = {};
let currentCow = '';
let currentFortune = '';

// Load data
async function loadData() {
  try {
    const [fortunesRes, cowsRes] = await Promise.all([
      fetch('fortunes.json'),
      fetch('cows.json')
    ]);
    fortunes = await fortunesRes.json();
    cows = await cowsRes.json();
    console.log(`Loaded ${fortunes.length} fortunes and ${Object.keys(cows).length} cows`);
    displayFortune();
  } catch (error) {
    console.error('Failed to load data:', error);
    document.getElementById('output').textContent = 'Failed to load Kamadhenu data. Are you running this from a server?';
  }
}

// Generate cowsay speech bubble
function generateCowsay(text, cowArt) {
  const maxWidth = 60; // Increased from 40 to match terminal better
  
  // Preserve existing line breaks, then wrap long lines
  const originalLines = text.split('\n');
  const lines = [];
  
  originalLines.forEach(line => {
    if (line.trim().length === 0) {
      // Preserve blank lines
      lines.push('');
    } else if (line.length <= maxWidth) {
      // Line is short enough, keep as-is
      lines.push(line.trim());
    } else {
      // Line is too long, wrap by words
      const words = line.trim().split(' ');
      let currentLine = '';
      
      words.forEach(word => {
        if ((currentLine + ' ' + word).trim().length > maxWidth) {
          if (currentLine) lines.push(currentLine.trim());
          currentLine = word;
        } else {
          currentLine += (currentLine ? ' ' : '') + word;
        }
      });
      if (currentLine) lines.push(currentLine.trim());
    }
  });

  // Build speech bubble
  const bubbleWidth = Math.max(...lines.map(l => l.length), 10);
  let bubble = ' ' + '_'.repeat(bubbleWidth + 2) + '\n';
  
  if (lines.length === 1) {
    bubble += `< ${lines[0].padEnd(bubbleWidth)} >\n`;
  } else {
    lines.forEach((line, i) => {
      const padded = line.padEnd(bubbleWidth);
      if (i === 0) {
        bubble += `/ ${padded} \\\n`;
      } else if (i === lines.length - 1) {
        bubble += `\\ ${padded} /\n`;
      } else {
        bubble += `| ${padded} |\n`;
      }
    });
  }
  
  bubble += ' ' + '-'.repeat(bubbleWidth + 2) + '\n';

  // Replace $thoughts with speech bubble connector
  const connector = '\\';
  const art = cowArt.replace(/\$thoughts/g, connector);
  
  return bubble + art;
}

// Random item from array
function random(arr) {
  return arr[Math.floor(Math.random() * arr.length)];
}

// Display a new fortune with typewriter effect
async function displayFortune() {
  if (fortunes.length === 0 || Object.keys(cows).length === 0) return;
  
  currentFortune = random(fortunes);
  const cowNames = Object.keys(cows);
  currentCow = random(cowNames);
  
  const output = generateCowsay(currentFortune, cows[currentCow]);
  const outputEl = document.getElementById('output');
  
  // Clear current output
  outputEl.textContent = '';
  
  // Split into lines and display one at a time
  const lines = output.split('\n');
  for (let i = 0; i < lines.length; i++) {
    outputEl.textContent += lines[i] + (i < lines.length - 1 ? '\n' : '');
    await new Promise(resolve => setTimeout(resolve, 40)); // 40ms delay between lines
  }
}

// Event listeners
document.getElementById('new-fortune').addEventListener('click', (e) => {
  e.preventDefault();
  displayFortune();
});

// Keyboard shortcuts (space or enter for new fortune)
document.addEventListener('keydown', (e) => {
  if ((e.code === 'Space' || e.code === 'Enter') && e.target === document.body) {
    e.preventDefault();
    displayFortune();
  }
});

// Load on start
loadData();
