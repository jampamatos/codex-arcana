const path = require('path');
const { app, BrowserWindow } = require('electron');

async function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
    },
  });

  // If we're in dev, load Vite's dev server
  if (!app.isPacked) {
    await win.loadURL('http://localhost:5173');
    win.webContents.openDevTools();
  } else {
    // Otherwise, load static index.html
    win.loadFile(path.join(__dirname, '../dist/index.html'));
  }
}

app.whenReady().then(createWindow);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit();
});
