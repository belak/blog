function systemDarkMode() {
  return window.matchMedia('(prefers-color-scheme: dark)').matches;
}

function getDarkMode() {
  const storedValue = localStorage.getItem('darkMode');
  if (storedValue !== undefined) {
    return storedValue !== 'false';
  }

  return systemDarkMode();
}

function toggleDarkMode() {
  const darkMode = getDarkMode();
  localStorage.setItem(
    'darkMode',
    darkMode ? 'false' : 'true',
  );

  updateDarkMode();
}

function updateDarkMode() {
  if (getDarkMode()) {
    document.body.classList.remove('light-mode');
    document.body.classList.add('dark-mode');
  } else {
    document.body.classList.add('light-mode');
    document.body.classList.remove('dark-mode');
  }
}

function clearDarkMode() {
  localStorage.removeItem('darkMode');
  updateDarkMode();
}

window.addEventListener('storage', updateDarkMode);
updateDarkMode();

let element = document.getElementById('dark-mode-toggle');
element.style.display = 'block';

// const mediaMatch = window.matchMedia('(prefers-color-scheme: dark)');
// mediaWatch.addEventListener(maybeToggleDarkMode)
