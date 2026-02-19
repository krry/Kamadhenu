#!/usr/bin/env node
// Build JSON data from fortunes and cows for web app

const fs = require('fs');
const path = require('path');

const projectRoot = path.join(__dirname, '..');
const webDir = path.join(projectRoot, 'web', 'public');
const fortunesDir = path.join(projectRoot, 'fortunes');
const cowsDir = path.join(projectRoot, 'cows');

// Create web/public if it doesn't exist
if (!fs.existsSync(webDir)) {
  fs.mkdirSync(webDir, { recursive: true });
}

console.log('🐄 Building Kamadhenu web data...\n');

// Parse fortunes
const fortunes = [];
const fortuneFiles = fs.readdirSync(fortunesDir).filter(f => !f.endsWith('.dat'));

fortuneFiles.forEach(file => {
  const content = fs.readFileSync(path.join(fortunesDir, file), 'utf8');
  const entries = content.split('%').map(f => f.trim()).filter(f => f.length > 0);
  fortunes.push(...entries);
});

fs.writeFileSync(
  path.join(webDir, 'fortunes.json'),
  JSON.stringify(fortunes, null, 2)
);
console.log(`✅ Fortunes parsed: ${fortunes.length} fortunes`);

// Parse cows
const cows = {};
const cowFiles = fs.readdirSync(cowsDir).filter(f => f.endsWith('.cow'));

cowFiles.forEach(file => {
  const cowName = path.basename(file, '.cow');
  const content = fs.readFileSync(path.join(cowsDir, file), 'utf8');
  
  // Extract ASCII art between <<EOC and EOC
  const match = content.match(/\$the_cow = <<EOC;\n([\s\S]*?)\nEOC/);
  if (match && match[1]) {
    cows[cowName] = match[1];
  }
});

fs.writeFileSync(
  path.join(webDir, 'cows.json'),
  JSON.stringify(cows, null, 2)
);
console.log(`✅ Cows parsed: ${Object.keys(cows).length} cows`);

console.log(`\n🎉 Build complete! Data in ${webDir}`);
