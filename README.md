## (WIP) An API for clipping YouTube videos.

### Features:
- Accepts YouTube URLs and returns downloadable media
- Can return in several formats (including audio only)
- Allows trimming start/end times and adjusting volume 

### Stack:
- Go: Core logic and concurrent job queue
- Gin: REST API framework with CORS support
- Redis: Job tracking and state storage
- Docker: Containerization 

### Planned:
- Frontend written in Svelte
- Tweaks to get "production ready"

Uses the `yt-dlp` and `ffmpeg` libraries.
