import { configureColorSchemeToggle } from './colorscheme.js';
import { configureHamburgerMenuEvents } from './hamburgermenu.js';

function main() {
  configureHamburgerMenuEvents();
  configureColorSchemeToggle();
}

document.addEventListener('DOMContentLoaded', main, false);
