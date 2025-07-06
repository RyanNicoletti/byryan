document.addEventListener("DOMContentLoaded", function () {
  const navLinks = document.querySelectorAll("nav a");
  const currentPath = window.location.pathname;

  navLinks.forEach((link) => {
    if (currentPath === link.getAttribute("href")) {
      link.classList.add("active");
    }
  });
});
