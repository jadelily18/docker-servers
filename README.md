# Docker Servers

This repo serves as a place for all my personal Minecraft server "templates" as docker-compose files and assets.

To use any of these templates, you'll need [Docker](https://github.com/docker/docker-install), [Python](https://www.python.org/downloads/), and [pip](https://pypi.org/project/pip/).

## Installation

Install Python dependencies using pip:

```bash
pip install requests wget InquirerPy
```

> You can also use `pip uninstall requests wget InquirerPy` after installing if you don't plan on using this on the same machine again!

Then, run the `install_template.py` script with curl:

```bash
curl https://raw.githubusercontent.com/jadelily18/docker-servers/main/install_template.py | python3 -
```

Now you can move to your newly-created directory and run your server with Docker compose using Docker Compose:

```bash
cd [server-directory]
docker compose up
```

<details>
  <summary>Side Note</summary>
  i hate managing python dependencies so, <i><b>so much.</b></i>
</details>
