# gopushbullet
Just a simple wrapper of go-pushbullet with some settings set in a configuration yaml

# Dependency
* [pushbullet]("github.com/xconstruct/go-pushbullet")

# Usage
* set you API key, channel tag, etc. in the configuration `yaml` file `config/pushbullet.yml`
* create a push bullet notifier by the `New("YOUR CONFIG PATH")` command
* push `Notify("title", "content")`
