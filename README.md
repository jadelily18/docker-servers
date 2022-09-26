# Docker Servers

This repo serves as a place for all my personal Minecraft server "templates" as docker-compose files and assets.

To use any of these templates, you'll need [Docker](https://github.com/docker/docker-install) and [Python](https://www.python.org/downloads/).

## Installation

Install Python dependencies using the following bash command:

```bash
pip install -r $(curl https://raw.githubusercontent.com/jadelily18/docker-servers/main/requirements.txt)
```

> You'll probably get a couple errors that I'm too lazy to fix, but they shouldn't affect anything.

Then, run the `install_template.py` script with curl:

```bash
curl https://raw.githubusercontent.com/jadelily18/docker-servers/main/install_template.py | python3 -
```

Now you can move to your newly-created directory and run your server with Docker compose using Docker Compose:

```bash
cd [server-directory]
docker compose up
```
