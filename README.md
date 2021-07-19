# Pingbot - Backend

## Install Pre-Compile binary

* Linux amd64
  * [Download](https://github.com/MedzikUser/go-pingbot/releases) latest version
  * Unpack file `tar xzf pingbot_{VERSION}_linux_amd64.tar.gz`
  * Create an .env and config.yml file and complete according to .env.schema and config.schema.yml
  * Add permissions `chmod +rwx pingbot.out`
  * Run binary `./pingbot.out`

## Disable Automatic Updates

* Add Option `-u=false` or `--update=false`
