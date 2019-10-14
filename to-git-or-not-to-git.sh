#! /bin/bash
curl https://api.github.com/users/DanabekKR | grep '"id"'| cut -d : -f 2 | cut -d , -f1 | cut -d "" -f2
