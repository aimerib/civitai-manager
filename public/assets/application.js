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

function checkMediaType(url, callback) {
  const img = new Image();
  img.onload = function () {
    callback("image");
  };
  img.onerror = function () {
    const video = document.createElement("video");
    video.onloadedmetadata = function () {
      callback("video");
    };
    video.onerror = function () {
      callback("unknown");
    };
    video.src = url;
  };
  img.src = url;
}

function loadImage(url) {
  return fetch(url)
    .then((response) => response.blob())
    .then((blob) => URL.createObjectURL(blob))
    .catch((error) => console.error("Error loading image:", error));
}

function setupMediaContainers() {
  const containers = document.querySelectorAll(".media-container");
  containers.forEach((container) => {
    const url = container.dataset.src;
    const alt = container.dataset.alt;

    loadImage(url).then((objectUrl) => {
      console.log("Loaded image:", objectUrl);
      const img = new Image();
      img.onload = function () {
        console.log("Image loaded successfully:", objectUrl);
        container.innerHTML = `<img src="${objectUrl}" alt="${alt}" class="w-full h-48 object-cover" style="object-position: left -20px">`;
        console.log("Image element added to container:", container.innerHTML);
      };
      img.onerror = function () {
        const video = document.createElement("video");
        video.onloadedmetadata = function () {
          container.innerHTML = `
            <video class="w-full h-48 object-cover" controls style="object-position: left -20px">
              <source src="${url}" type="video/mp4">
              Your browser does not support the video tag.
            </video>
          `;
        };
        video.onerror = function () {
          container.innerHTML = `<p>Unsupported media type</p>`;
        };
        video.src = url;
      };
      img.src = objectUrl;
    });
  });
}

if ("serviceWorker" in navigator) {
  window.addEventListener("load", function () {
    navigator.serviceWorker.register("/sw.js").then(
      function (registration) {
        console.log(
          "ServiceWorker registration successful with scope: ",
          registration.scope
        );
      },
      function (err) {
        console.log("ServiceWorker registration failed: ", err);
      }
    );
  });
}

htmx.onLoad(setupMediaContainers);
htmx.onLoad(setupSidebar);
