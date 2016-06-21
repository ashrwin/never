# never


never is a CLI tool that builds and (re)starts your web application whenever there is a change in the source files.


## Installation

    go get github.com/ashrko619/never

## Usage

    cd /path/to/myapp

Start never:

    never [args]

Never will watch for file events, and every time you create/modifiy/delete a file it will build and restart the app. Arguments passed to `never` will be forwarded to your app.


## Configuration

`never` looks for a `'never.conf.json'` file in the current working directory for additional configuration. Here is a sample format for the config file

```javascript
{
	"ignoredFolders" :["js","css"],       // folders to ignore (recursively)
	"ignoredExtensions" : [".html",".js"] // extensions to ignore
}
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


