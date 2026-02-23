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

function escapeHtml(str) {
  return str
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#039;');
}

async function renderRainbowOutput(output) {
  const outputEl = document.getElementById('output');
  outputEl.innerHTML = '';

  const lines = output.split('\n');
  for (let i = 0; i < lines.length; i++) {
    const line = document.createElement('span');
    line.className = `cowsay-line rainbow-${i % 7}`;
    line.innerHTML = escapeHtml(lines[i]);
    outputEl.appendChild(line);
    await new Promise(resolve => setTimeout(resolve, 35));
  }
}

// Display a new fortune with typewriter + rainbow effect
async function displayFortune() {
  if (fortunes.length === 0 || Object.keys(cows).length === 0) return;

  currentFortune = random(fortunes);
  const cowNames = Object.keys(cows);
  currentCow = random(cowNames);

  const output = generateCowsay(currentFortune, cows[currentCow]);
  await renderRainbowOutput(output);
}

async function displayAliasFortune() {
  if (Object.keys(cows).length === 0) return;

  const aliasInput = window.prompt('Put words in Kamadhenu\'s mouth:', currentFortune || 'Do you love me?');
  if (aliasInput === null) return;

  const trimmed = aliasInput.trim();
  if (!trimmed) {
    return displayFortune();
  }

  const cowNames = Object.keys(cows);
  currentCow = random(cowNames);
  currentFortune = trimmed;

  const output = generateCowsay(currentFortune, cows[currentCow]);
  await renderRainbowOutput(output);
}

// Event listeners
const newFortuneBtn = document.getElementById('new-fortune');
newFortuneBtn.addEventListener('click', (e) => {
  e.preventDefault();
  displayFortune();
});

// Tap on oracle to refresh (click handles taps)
const oracleWrapper = document.querySelector('.oracle-wrapper');
oracleWrapper.addEventListener('click', (e) => {
  // Don't trigger if clicking links/footer or selecting text
  if (e.target.closest('footer') || e.target.closest('a') || window.getSelection().toString()) return;
  displayFortune();
});

// Swipe detection variables
let touchStartX = 0;
let touchStartY = 0;

// Swipe logic restricted to oracle area
oracleWrapper.addEventListener('touchstart', (e) => {
  touchStartX = e.changedTouches[0].screenX;
  touchStartY = e.changedTouches[0].screenY;
}, { passive: true });

oracleWrapper.addEventListener('touchend', (e) => {
  const touchEndX = e.changedTouches[0].screenX;
  const touchEndY = e.changedTouches[0].screenY;
  
  const diffX = Math.abs(touchEndX - touchStartX);
  const diffY = Math.abs(touchEndY - touchStartY);
  
  // Horizontal swipe trigger (prevent vertical scroll from triggering)
  if (diffX > 50 && diffX > diffY) {
    displayFortune();
  }
}, { passive: true });

// Keyboard shortcuts (space or enter for new fortune)
document.addEventListener('keydown', (e) => {
  if ((e.code === 'Space' || e.code === 'Enter') && e.target === document.body) {
    e.preventDefault();
    displayFortune();
  }
});

// Load on start
loadData();
