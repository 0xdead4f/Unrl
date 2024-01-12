![unrl logo](unrl.png)

# UNRL
---

`Unrl` or Unique (instead of Uniform because i have no idea what what to name it and it's 1:00 AM and im very sleepy) Resource Locator Eliminating parameter value and then remove duplicates url from stdin list input.

## Installation
```
go install github.com/0xdead4f/Unrl@latest
```

## Usage
```
> cat url.txt
https://example.com/test?a=aaaa&b=bbbbb
https://example.com/test?a=cccc&b=ddddd
https://example.com/test?a=zzzz&b=xxxxx
https://example.com/diff?a=aaaa&b=bbbbb
https://example.com/diff?a=aaaa&b=ddddd
https://example.com/diff?a=aaaa&b=ddddd&c=zzzz
```
Use ```unrl``` from stdin :
```
> cat url.txt | unrl

https://example.com/diff?a=&b=&c=
https://example.com/test?a=&b=
```

## Personal Use
Chaining with other tools to find reflected element in websites :
```
subfinder -d example.com | gau | unrl | kxss | grep -v "\[\]" | tee result.txt
```
## Next Objective
The next objective of this project is eliminating similiar url with different id identifier in the path. For example :
```
https://example.com/blog/123123/?s=query
https://example.com/blog/321312/?s=query
```
The url should be marked same because it's basically same url and have same functionality. If you guys know how to effectively do such an method, please do a pull request !
