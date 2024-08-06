htmx.config.globalViewTransitions = true;
// htmx.logAll();

function setupSidebar() {
  const sidebar = document.getElementById("sidebar");
  const sidebarTrigger = document.getElementById("sidebar-trigger");

  if (sidebar && sidebarTrigger) {
    if (!sidebarTrigger.getAttribute("data-has-mouseenter")) {
      sidebarTrigger.addEventListener("mouseenter", () => {
        sidebar.classList.remove("-translate-x-full");
      });
      sidebarTrigger.setAttribute("data-has-mouseenter", "true");
    }

    if (!sidebar.getAttribute("data-has-mouseleave")) {
      sidebar.addEventListener("mouseleave", () => {
        sidebar.classList.add("-translate-x-full");
      });
      sidebar.setAttribute("data-has-mouseleave", "true");
    }
  }
}

htmx.onLoad(setupSidebar);

document.addEventListener("alpine:init", () => {
  Alpine.store("sidebar", {
    open: false,
    toggle() {
      this.open = !this.open;
    },
  });
});

document.body.addEventListener("htmx:afterSettle", (event) => {
  Alpine.initTree(document.body);
});

document.body.addEventListener("htmx:historyRestore", (event) => {
  Alpine.initTree(document.body);
});
