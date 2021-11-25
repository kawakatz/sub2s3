# sub2s3
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/sub2s3/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/kawakatz/sub2s3"><img src="https://goreportcard.com/badge/github.com/kawakatz/sub2s3"></a>
<a href="https://www.codefactor.io/repository/github/kawakatz/sub2s3/badge"><img src="https://www.codefactor.io/repository/github/kawakatz/sub2s3/badge"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

Check if subdomains are on S3 or not.<br>
sub2s3 will sleep 1 second after each GET request to AWS servers.<br>
## installation
```sh
➜  ~ go install -v github.com/kawakatz/sub2s3@latest
```

## usage
```sh
➜  ~ cat subdomains.txt | sub2s3
```
