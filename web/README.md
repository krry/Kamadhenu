# Kamadhenu Web

An udderly absurd web oracle. Random fortunes delivered by ASCII cows with rainbow gradients.

## 🌐 Live Site

Deploy this to Netlify and watch the wisdom flow.

## 🏗️ Build

The build script parses all fortunes and cows into JSON:

```bash
node ../scripts/build-web-data.js
```

This generates:
- `public/fortunes.json` (13,443 fortunes)
- `public/cows.json` (105 ASCII cows)

## 🚀 Deploy to Netlify

1. Connect this repo to Netlify
2. Set build settings:
   - **Build command:** `node ../scripts/build-web-data.js`
   - **Publish directory:** `public`
   - **Base directory:** `web`

Or use the Netlify CLI:

```bash
cd web
netlify deploy --prod
```

## 💻 Local Development

Serve the `public` directory with any static server:

```bash
# Python
cd public && python3 -m http.server 8000

# Node
npx serve public

# Bun
bunx serve public
```

## ✨ Features

- **13,443 fortunes** from classic fortune files
- **105 ASCII cows** (Star Wars, banana, dragon, and more)
- **Rainbow mode** (lolcat-style gradient)
- **Keyboard shortcut:** Press `Space` for new fortune
- **Mobile-friendly** responsive design

## 🎨 Customization

Edit `public/style.css` to change colors, gradients, or layout.

The cowsay speech bubble generation is in `public/app.js` (`generateCowsay` function).

## 📦 Data Structure

**fortunes.json:**
```json
[
  "A day for firm decisions!!!!! Or is it?",
  "A few hours grace before the madness begins again.",
  ...
]
```

**cows.json:**
```json
{
  "banana": "       \\n        \\n\\n     \".           ,#  \\n     \\\\ `-._____,-'=/\\n  ____`._ ----- _,'_____PhS\\n         `-----'\\n",
  "dragon": "..."
}
```

## 🐄 About Kamadhenu

Kamadhenu is the wish-granting divine cow from Hindu mythology.

This web app brings the CLI fortune/cowsay experience to the browser with:
- Random fortune selection
- Random cow selection
- Cowsay-compatible speech bubble generation
- Rainbow gradient effects (inspired by lolcat)

Built with vanilla HTML, CSS, and JavaScript. No frameworks. No build step (except data parsing). Just pure web goodness.

---

**Zuckle at the zipple of wisdom** 🐄✨
