# Docker Servers

This repo serves as a place for all my personal Minecraft server "templates" as docker-compose files and assets.

To use any of these templates, you'll need [Docker](https://github.com/docker/docker-install), [Python](https://www.python.org/downloads/), and [pip](https://pypi.org/project/pip/).

## Installation

### Linux/OSX

Download the respective file from the `build` directory. Using `cURL` for the Linux file:

```bash
curl -sL https://github.com/jadelily18/docker-servers/blob/main/build/get-template-linux\?raw\=true --output get-template-linux
```

Then give it executable permissions with `chmod +x` and then run it:

```bash
chmod +x get-template-linux && ./get-template-linux
```

Now you can move to your newly-created directory and run your server with Docker compose using Docker Compose:

```bash
cd [server-directory]
docker compose up
```

### Windows

Download the `get-template-windows.exe` file from the `build` directory and run in in Command Prompt or Powershell:

```powershell
curl -sL https://github.com/jadelily18/docker-servers/blob/main/build/get-template-linux\?raw\=true --output get-template-windows.exe
get-template-windows.exe
```

Now you can move to your newly-created directory and run your server with Docker compose using Docker Compose:

```bash
cd [server-directory]
docker compose up
```
