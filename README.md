# DRACO: Dynamic Response and Command Optimization

**Dynamic Response:**
- Highly-adaptable suggestions.
- Fits 99.9% of use cases within repetition filled linux workflows. 

**Command Optimization:**
- Streamlined menu system for quick access to common commands.
- Reduces repetitive typing, enhancing efficiency.

**AR/XR Integration:**
- Optimized for wide screens and minimal inputs.
- Seamless use with augmented and extended reality devices.
- Built with AR in mind, large viewing angle through the UI, minimal key inputs. 

**Flexible Configuration:**
- Easily customizable via config.json.
- Auto-generates logs in /usr/var/logs/draco* (modifiable path).

# Strategic Project Setup and Deployment
- Our project structure is designed with clarity and scalability in mind. Simply place your config.json at the root of the project to get started. Run the program with root permissions to auto-generate a logging directory at /usr/var/logs/draco*. While this location is unique it is chosen for optimal system integration (randomly typed as i was coding), customization is at your fingertips—adjust it as needed to fit your environment.


Anyway file structure looks like this for the project:

```
draco/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── model/
│   │   └── model.go
│   ├── ui/
│   │   └── ui.go
└── go.mod
```

### INSTALLATION

1. git clone this repository
2. go mod tidy
3. go build -o draco ./cmd/main.go
4. binary should be created for you, move this to /usr/bin as usual if you want it accross the system. 

Note: Minimal dependencies should be needed, the program needs to be ran as root to ensure commands in there are executed + the location of the log files, chown it if you want.


### PROGRAM CONFIG METHOD AND COLOR PALETTES:

Here is the overview on configuring, its straightforward if you know english. Or JSON for that matter. Just add in more menus how your use case fits, more programs in there etc :)

```

  "menus": {
    "MENU NAME": [
      {

      Like this, add in what you want, either under pre build menus or make a new menu!

        "name": "COMMAND",
        "description": "COMMAND DESCRIPTION",
        "arguments": ["-ARG", "-ARG", "-ARG"],
        "paths": []

      Paths are a work in progress still, I am still build in better variable hadling accross programs and menus, etc

Default color palette

  "colorPalette": {
    "cursorColor": "#802e2e",
    "selectedColor": "8f2424",
    "textColor": "#F3f6f4",
    "backgroundColor": "#333333"

Cyberpunk/Retro Vibes

  "colorPalette": {
    "cursorColor": "#ff69b4",        // Hot Pink
    "selectedColor": "#87CEEB",      // Sky Blue
    "textColor": "#e0ffff",          // Light Cyan
    "backgroundColor": "#2d2d44"     // Dark Indigo

Neon Noir

  "colorPalette": {
    "cursorColor": "#ff4500",        // Neon Orange-Red
    "selectedColor": "#00ffcc",      // Neon Aqua
    "textColor": "#d1c4e9",          // Lavender Mist
    "backgroundColor": "#0d0d1a"     // Deep Midnight Blue

Green and Red (Hacker Matrix Vibes)

  "colorPalette": {
    "cursorColor": "#00ff00",        // Neon Green
    "selectedColor": "#ff0000",      // Bright Red
    "textColor": "#00ff00",          // Neon Green
    "backgroundColor": "#000000"     // Black

Synthwave (Retro 80s)

  "colorPalette": {
    "cursorColor": "#ff00ff",        // Fuchsia
    "selectedColor": "#00ffff",      // Cyan
    "textColor": "#ffcc00",          // Neon Yellow
    "backgroundColor": "#000033"     // Deep Navy Blue

Desert Sunrise (Warm Tones)

  "colorPalette": {
    "cursorColor": "#ff7f50",        // Coral
    "selectedColor": "#ffd700",      // Gold
    "textColor": "#fff5ee",          // Seashell
    "backgroundColor": "#8b4513"     // Saddle Brown

Cool Blues (Tech Minimalist)

  "colorPalette": {
    "cursorColor": "#4682b4",        // Steel Blue
    "selectedColor": "#5f9ea0",      // Cadet Blue
    "textColor": "#e0ffff",          // Light Cyan
    "backgroundColor": "#2c3e50"     // Midnight Blue


```

### License = GPLv3


https://www.gnu.org/licenses/gpl-3.0.en.html


The Foundations of the GPL
Nobody should be restricted by the software they use. There are four freedoms that every user should have:


- the freedom to use the software for any purpose,
- the freedom to change the software to suit your needs,
- the freedom to share the software with your friends and neighbors, and
- the freedom to share the changes you make.

When a program offers users all of these freedoms, we call it **free software**.

Developers who write software can release it under the terms of the GNU GPL. When they do, it will be free software and stay free software, no matter who changes or distributes the program. We call this copyleft: the software is copyrighted, but instead of using those rights to restrict users like proprietary software does, we use them to ensure that every user has freedom.


