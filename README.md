Github repo utility
=======================
A github API command-line utility which fetches the most recent pull requests, or most recent releases, for a github repo. 

Example usage: 
-----------
```
./github-repo-utility -e pulls -i https://api.github.com/repos/mailchimp/mc-woocommerce
``` 
Build:
--------
```
go build -o github-repo-utility
```

Docker:
-----------
Build:
```
docker build -t github-repo-utility:latest .
```

Run:
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
 - better test coverage && add test fixtures
 - add "go test" step to github action workflow
 - edit/improve logging text
 - combine duplicate functions: listGithubReleases, listGithubPulls 
 - reconsider regex implementation
 - (optional) add -b [branch] switch
 - (optional) remove default repo URL
 - (optional) add metrics "hooks"
