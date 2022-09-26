#!/bin/bash

python3 -m pip install pipenv

pipenv shell
pipenv install requests inquirer wget

curl -s https://raw.githubusercontent.com/jadelily18/docker-servers/main/install_template.py -O ~/install_template.py
python3 ~/install_template.py

pipenv --rm
deactivate

