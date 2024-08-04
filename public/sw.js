const CACHE_NAME = 'media-cache-v9';

self.addEventListener('install', function(event) {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => {
      console.log('Cache opened:', CACHE_NAME);
    })
  );
});

self.addEventListener('fetch', function(event) {
  if (event.request.url.match(/\.(jpg|jpeg|png|gif|svg)$/)) {
    event.respondWith(
      caches.match(event.request)
        .then(function(response) {
          console.log('Fetching image:', event.request.url);
          if (response) {
            console.log('Cache hit for:', event.request.url);
            return response;
          }
          console.log('Cache miss for:', event.request.url);
          return fetch(event.request).then(function(networkResponse) {
            if (!networkResponse || networkResponse.status !== 200) {
              console.log('Network fetch failed for:', event.request.url);
              return networkResponse;
            }
            console.log('Network fetch succeeded for:', event.request.url);
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
            console.log('Deleting old cache:', cacheName);
            return caches.delete(cacheName);
          }
        })
      );
    })
  );
});
