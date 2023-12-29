export function configureHamburgerMenuEvents() {
  document.getElementById('hamburger-open-button').addEventListener('click', function () {
    document.querySelector('.page-header').setAttribute('data-hamburger-state', 'open');
  });

  document.getElementById('hamburger-close-button').addEventListener('click', function () {
    document.querySelector('.page-header').setAttribute('data-hamburger-state', 'closed');
  });
}
