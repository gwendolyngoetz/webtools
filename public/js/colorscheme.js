export function configureColorSchemeToggle() {
  let lightModeState = false;
  const toggleButton = document.getElementById('color-scheme-toggle');
  const useLight = window.matchMedia('(prefers-color-scheme: light)');

  function toggleLightMode(state) {
    document.documentElement.classList.toggle('light-mode', state);
    lightModeState = state;
  }

  function setLightModeLocalStorage(state) {
    localStorage.setItem('light-mode', state);
  }

  toggleLightMode(localStorage.getItem('light-mode') == 'true');
  toggleButton.checked = !(localStorage.getItem('light-mode') == 'true');

  useLight.addEventListener('change', (event) => toggleLightMode(event.matches));

  toggleButton.addEventListener('click', () => {
    lightModeState = !lightModeState;
    toggleLightMode(lightModeState);
    setLightModeLocalStorage(lightModeState);
  });
}
