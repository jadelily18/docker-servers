import os

if os.path.isfile("docker-compose.yml"):
    print("Running server:")
    os.system("docker compose up -d")
    print("Use 'docker attach [CONTAINER NAME]' to attach.")
else:
    print("No 'docker-compose.yml' file found. Downloading file:")
    os.system("wget https://raw.githubusercontent.com/jadelily18/docker-servers/main/docker-compose/jade-kris/docker-compose.yml")
    print("Running server:")
    os.system("docker compose up -d")
    print("Use 'docker attach [CONTAINER NAME]' to attach.")
