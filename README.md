# PROMETHEUS EXPORTER BOILERPLATE

A working example to build prometheus exporter

## Features

* Cobra and Viper to handle environment variables and cli flags
* standard library logger with `log/slog`
* Working Dockerfile example to build multi arch docker images

## How to get started

Just copy the project and tweak it to your needs. 

## Structure

*cmd/* contains the CLI aspect of the exporter, loads the CLI flags, initializes the settings, the exporter and starts the web server.
*internal/* contains the internal packages for the different components of the exporter, this packages are only to be used by the project and cannot be imported by other projects
*internal/collectors* contains the exporter logic
*internal/httpServer* contains the http server responsible for serving the metrics
*internal/slogLogger* contains the logger initialization logic

# Contributions

Improvements and suggestions are always welcome, feel free to check for any open issues, open a new Issue or Pull Request

If you like this project and want to support / contribute in a different way you can always: 

<a href="https://www.buymeacoffee.com/coolapso" target="_blank">
  <img src="https://cdn.buymeacoffee.com/buttons/default-yellow.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" />
</a>

# Projects built using this example

* [prometheus-twitch-exporter](https://github.com/coolapso/prometheus-twitch-exporter)
