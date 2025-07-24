## A tool for downloading videos from the internet.

### Key Features:
- Accepts URLs and returns downloadable media
- Can convert to several formats (including audio only)
- Allows trimming and volume adustment

### Technical Details:
- Fully containerized and reproducable in any environment
- Ready for local use or easy deployment to a VPS
- Centralized config with `docker-compose`

### Stack: 
- Gin: REST API framework
- Go: Queue-based job processing
- Redis: Job tracking and state storage
- Docker: Containerization
- Nginx: Reverse proxy

### Notes:
- A fully-featured SvelteKit frontend for this project: (clip-fe)[https://github.com/thomascpowell/clip-fe].
- The frontend supports self-hosting, since I don't currently provide a hosted backend
