const CACHE_NAME = 'media-cache-v10';

self.addEventListener('install', function(event) {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => {
    })
  );
});

self.addEventListener('fetch', function(event) {
  if (event.request.url.match(/\.(jpg|jpeg|png|gif|svg)$/)) {
    event.respondWith(
      caches.match(event.request)
        .then(function(response) {
          if (response) {
            return response;
          }
          return fetch(event.request).then(function(networkResponse) {
            if (!networkResponse || networkResponse.status !== 200) {
              return networkResponse;
            }
            var responseToCache = networkResponse.clone();
            caches.open(CACHE_NAME)
              .then(function(cache) {
                cache.put(event.request, responseToCache)
                  .then(() => console.log('Cached:', event.request.url))
                  .catch(err => console.error('Caching failed for:', event.request.url, err));
              });
            return networkResponse;
          });
        })
        .catch(function(error) {
          console.error('Error in fetch handler:', error);
          return new Response('Image not available', {status: 404, statusText: 'Not Found'});
        })
    );
  }
});

self.addEventListener('activate', function(event) {
  event.waitUntil(
    caches.keys().then(function(cacheNames) {
      return Promise.all(
        cacheNames.map(function(cacheName) {
          if (cacheName !== CACHE_NAME) {
            return caches.delete(cacheName);
          }
        })
      );
    })
  );
});
