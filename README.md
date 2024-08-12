# DRACO

## Linux Helper and Repetition Eliminator

You repeating yourself in the terminal? Speed things up with this handy menu :)
This is als ooptimised for augmented reality with wide screens, minimal inputs etc. Plug in your XR device and have at it.

### FILE STRUCT:


File structure index of the go files etc etc, place your config.json at the projects root, Run the program with root perms to create a logging directory in /usr/var/logs/draco* "Very Different I know!"
Oh well change it if you dont like it this is where im putting it for now :)

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
4. binary should be created for you, move this to /usr/bin as usual if you want it accross the system. Or keep it there idk. Im not your mother.

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


