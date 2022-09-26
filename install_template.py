import os
import requests
import inquirer
import wget


templates = requests.get(
    "https://api.github.com/repos/jadelily18/docker-servers/contents/docker-compose").json()
template_names = []

for i in templates:
    template_names.append(i["name"])

questions = [
    inquirer.Text(
        "serverName", "What would you like to call your server directory?", default="mc-server"),
    inquirer.List(
        'template', message="Select from project templates", choices=template_names)
]

answers = inquirer.prompt(questions)

server_dir_name = answers["serverName"]
repo_name = answers["template"]

docker_compose_url = requests.get(
    f"https://api.github.com/repos/jadelily18/docker-servers/contents/docker-compose/{repo_name}/docker-compose.yml").json()["download_url"]

if os.path.exists(os.path.join(server_dir_name, "docker-compose.yml")):
    print("\nDirectory already exists! Please retry with a different directory name.")
else:
    os.mkdir(server_dir_name)
    wget.download(docker_compose_url, out=os.path.join(
        server_dir_name, "docker-compose.yml"))
    print("\n\nYou can now run your server using `docker compose up`. If you don't already have docker installed, you can use its their convenience script at https://github.com/docker/docker-install.\n")
