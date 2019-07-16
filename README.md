### Development
install govendor first `https_proxy=http://127.0.0.1:1087 go get -u github.com/kardianos/govendor`
```
# Add existing GOPATH files to vendor.
govendor add +external
```

### config format
```json
{
  "corpId": "YOUR CORP ID",
  "corpSecret": "APP SECRET",
  "agentId": APP AGENT ID
}
```

### Test
See [example](https://github.com/dongfg/notify/tree/master/example) folder
Example List:
- [x] Text
- [x] Image
- [x] Voice
- [x] Video
- [x] File
- [x] TextCard
- [x] News
- [x] MpNews
- [x] Markdown
