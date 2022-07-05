Currency converter
=======================
A github API command-line utility which fetches the most recent pull requests, or most recent releases, for a github repo. 

Example usage: 
-----------
```
go run main.go -e pulls -i https://api.github.com/repos/mailchimp/mc-woocommerce
``` 

Docker:
-----------
Build a docker image with:
```
docker build -t github-repo-utility:latest .
```

Run the docker image

```
docker run --rm -it github-repo-utility:latest 
```
Help
--------
```
docker run --rm -it github-repo-utility:latest -h
```

To do
------------
 - add unit tests to get better coverage
 - add "go test" step to github action workflow
 - edit/improve logging text
 - reconsider regex implementation
 - (optional) add -b [branch] switch
 - (optional) remove default repo URL
 - (optional) add metrics 

