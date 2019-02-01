# famousPersonRadar
Use wiki to detect that whether this is a famous Chinese name

# How to use
- Put your name list in personList.txt.(A name per line)
- Run the codes by using: go run famousPersonRadar.go(You have to install golang first)
- You will get the output in output.txt. 1 means that this is a famous person, 0 means unfamous

# Principle
Get the http status code from https://zh.wikipedia.org/wiki/XXX

200 - famous, 404 - unfamous

# Tips
This tool is for Chinese person names only
