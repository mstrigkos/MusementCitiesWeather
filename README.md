# Musement Cities Weather

Musement Cities Weather is a simple CLI application written in Go. The application fetches the cities that [Musement](https://www.musement.com/us/) has available activities and prints to STDOUT a 2-day weather forecast. The weather info is fetched from the [WeatherAPI](https://www.weatherapi.com/).

## Installation

Use the [git](https://git-scm.com/) version control in order to download the project on your system.

```bash
git clone https://github.com/mstrigkos/MusementCitiesWeather
```
The next step is to build the [Docker](https://www.docker.com/get-started) image and get the image ready to run the application.
```bash
cd MusementCitiesWeather
docker build --tag musementcitiesweather .
```
## Configuration (optional)
In order to update the parameters of the application you can edit the [config.json](https://github.com/mstrigkos/MusementCitiesWeather/blob/master/config.json) file.
```bash
{
    "musementAPIURL": "https://sandbox.musement.com/api/v3/cities", #The Musement API URL
    "logFile": "errors.log", #Error log
    "weatherAPIURL": "https://api.weatherapi.com/v1/forecast.json", #Weather API URL
    "weatherAPIKey": "30ba692deba345f19f9185654211411" #Weather API key
}
```

## Usage
Now the docker image can be run using the following command:
```bash
docker run musementcitiesweather 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Developed by Michael Strigkos

[MIT](https://github.com/mstrigkos/MusementCitiesWeather/blob/master/LICENSE)
