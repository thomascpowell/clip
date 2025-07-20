## (WIP) An API for clipping YouTube videos.

### Features:
- Accepts YouTube URLs and returns downloadable media
- Can return in several formats (including audio only)
- Allows trimming start/end times and adjusting volume 

### Technical Details:
- RESTful API built with Gin
- Concurrent job queue
- Redis for state storage
- Fully containerized

### Planned Features:
- Frontend written in Svelte
- Tweaks to get "production ready"
