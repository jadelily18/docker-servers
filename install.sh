#!/bin/bash

sudo apt -y install python3-pip

pip install requests inquirer wget

curl https://raw.githubusercontent.com/jadelily18/docker-servers/main/install_template.py | python3 -

