# Weather CLI Example

This is an example of the weather CLI based on the [WeatherAPI](https://www.weatherapi.com/) implemented with the [Go (golang)](https://go.dev/) programming language.

## Requirements for development

Go must be installed.

## Compile

```sh
go build -o weather-cli .
```

## Requirements for usage

API-Key from [WeatherAPI](https://www.weatherapi.com/) must be set as an environment variable `WEATHERAPI_KEY`.

## Usage

Show help

```sh
./weather-cli -h
```

Show weather for Berlin

```sh
./weather-cli -l Berlin
```
