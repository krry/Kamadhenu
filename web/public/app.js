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
  const maxWidth = 40;
  const words = text.split(' ');
  const lines = [];
  let currentLine = '';

  // Wrap text
  words.forEach(word => {
    if ((currentLine + ' ' + word).trim().length > maxWidth) {
      lines.push(currentLine.trim());
      currentLine = word;
    } else {
      currentLine += (currentLine ? ' ' : '') + word;
    }
  });
  if (currentLine) lines.push(currentLine.trim());

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
  const connector = lines.length === 1 ? '\\' : '\\';
  const art = cowArt.replace(/\$thoughts/g, connector);
  
  return bubble + art;
}

// Random item from array
function random(arr) {
  return arr[Math.floor(Math.random() * arr.length)];
}

// Display a new fortune
function displayFortune() {
  if (fortunes.length === 0 || Object.keys(cows).length === 0) return;
  
  currentFortune = random(fortunes);
  const cowNames = Object.keys(cows);
  currentCow = random(cowNames);
  
  const output = generateCowsay(currentFortune, cows[currentCow]);
  const outputEl = document.getElementById('output');
  outputEl.textContent = output;
  
  // Apply rainbow mode if enabled
  const rainbowMode = document.getElementById('rainbow-mode').checked;
  const h1 = document.querySelector('h1');
  
  if (rainbowMode) {
    outputEl.classList.add('rainbow');
    h1.classList.add('rainbow');
  } else {
    outputEl.classList.remove('rainbow');
    h1.classList.remove('rainbow');
  }
}

// Event listeners
document.getElementById('new-fortune').addEventListener('click', displayFortune);
document.getElementById('rainbow-mode').addEventListener('change', displayFortune);

// Keyboard shortcut (space bar for new fortune)
document.addEventListener('keydown', (e) => {
  if (e.code === 'Space' && e.target === document.body) {
    e.preventDefault();
    displayFortune();
  }
});

// Load on start
loadData();
