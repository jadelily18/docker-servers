#!/bin/bash

sudo apt -y install python3-pip

pip install pipenv


curl -s https://raw.githubusercontent.com/jadelily18/docker-servers/main/install_template.py -O ~/install_template.py
pipenv run pip install requests InquirerPy wget
pipenv run python3 ~/install_template.py

pipenv --rm

