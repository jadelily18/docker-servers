import os

if os.path.isfile("docker-compose.yml"):
    print("Running server:")
    os.system("docker compose up")
else:
    print("No 'docker-compose.yml' file found. Downloading file:")
    os.system("wget https://raw.githubusercontent.com/jadelily18/docker-servers/main/docker-compose/jade-kris/docker-compose.yml")
    os.system("docker compose up")
