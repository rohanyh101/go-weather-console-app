# Go Weather Console API
#### Video Demo: <URL HERE>
#### Description:

## Overview

The Go Weather Console API is a command-line weather application built in Go that fetches real-time weather data and hourly forecasts from the WeatherAPI service. This project demonstrates the power of Go's simplicity and efficiency in creating practical command-line tools that provide users with essential weather information in an elegant, colorized terminal interface.

The application allows users to quickly check current weather conditions and hourly forecasts for any city worldwide, with intelligent color coding to highlight important weather events such as high chances of rain. By default, it shows weather information for Jamkhandi, but users can specify any city as a command-line argument to get localized weather data.

## Project Architecture and Files

### main.go
The `main.go` file serves as the heart of the application and contains all the core functionality. This file demonstrates several important Go programming concepts and best practices:

**Data Structures**: The file defines a comprehensive `Weather` struct that maps to the JSON response from the WeatherAPI. This struct uses Go's JSON tags to properly unmarshal the API response, showcasing Go's excellent JSON handling capabilities. The nested structure includes Location, Current weather conditions, and Forecast data with hourly breakdowns.

**HTTP Client Implementation**: The application uses Go's built-in `net/http` package to make API requests. The code properly handles HTTP responses, including error checking for non-200 status codes and proper resource cleanup using `defer resp.Body.Close()`.

**Environment Variable Management**: Security-conscious design is implemented by storing the API key in environment variables rather than hardcoding sensitive information. The application uses the `godotenv` package to load environment variables from a `.env` file.

**Command-line Argument Processing**: The application accepts command-line arguments to specify different cities, defaulting to "Jamkhandi" if no argument is provided. This demonstrates Go's `os.Args` slice handling.

**Time Manipulation**: The code converts Unix timestamps to Go's `time.Time` objects and formats them for display, showing proficiency with Go's time package.

**Conditional Logic and Data Filtering**: The hourly forecast display includes intelligent filtering to only show future hours (skipping past hours) and color-codes entries based on rain probability.

### go.mod and go.sum
The `go.mod` file defines the module path and Go version requirements, while `go.sum` ensures dependency integrity. This project uses several external dependencies:

- `github.com/fatih/color`: Provides cross-platform colored terminal output
- `github.com/joho/godotenv`: Enables loading environment variables from `.env` files
- Additional transitive dependencies for color rendering on different platforms

### weather.json
This file appears to contain sample weather data, likely used for testing or development purposes. It demonstrates the structure of the WeatherAPI response and helps in understanding the data format that the application processes.

### .gitignore
The `.gitignore` file properly excludes the `.env` file from version control, ensuring that sensitive API keys are not accidentally committed to the repository. This follows security best practices for handling API credentials.

## Design Decisions and Implementation Choices

**API Choice**: I chose WeatherAPI.com as the data source because it provides comprehensive weather data including hourly forecasts, has a generous free tier, and returns well-structured JSON responses that are easy to parse in Go.

**Color Coding Strategy**: The application uses color coding to enhance user experience - red for high rain probability (≥40%) and cyan for normal conditions. This visual distinction helps users quickly identify when they might need to carry an umbrella or plan indoor activities.

**Error Handling**: The application implements robust error handling at multiple levels - HTTP request failures, API response errors, JSON unmarshaling errors, and environment variable validation. This ensures the application fails gracefully with informative error messages.

**Data Filtering Logic**: The decision to filter out past hours from the forecast was made to provide only actionable information. Users typically want to know what weather to expect in the coming hours, not what already happened.

**Default City Selection**: Hubli was chosen as the default city, likely representing the developer's location or a meaningful place. This provides a sensible default while maintaining flexibility for users to specify other locations.

**Struct Design**: The Weather struct closely mirrors the API response structure but only includes the fields needed by the application. This selective approach reduces memory usage and makes the code more maintainable.

## Usage and Features

The application provides several key features:

1. **Current Weather Display**: Shows the current temperature, location, and weather conditions in a clear format.

2. **Hourly Forecast**: Displays upcoming hourly weather data including temperature, rain probability, and conditions.

3. **Visual Indicators**: Uses color coding to highlight important weather events, particularly high chances of rain.

4. **Flexible Location Input**: Accepts any city name as a command-line argument for global weather information.

5. **Time-Aware Filtering**: Automatically filters out past hours to show only relevant upcoming weather.

## Technical Achievements

This project demonstrates proficiency in several Go programming concepts:

- HTTP client programming and API integration
- JSON parsing and data structure mapping  
- Environment variable management and security practices
- Command-line argument processing
- Time manipulation and formatting
- Error handling and logging
- External package management and dependency handling
- Cross-platform terminal color output
- Struct design and data modeling

The application showcases Go's strengths in creating efficient, readable, and maintainable command-line tools. The code is structured in a single file for simplicity while maintaining good separation of concerns through clear variable naming and logical flow.

## Future Enhancements

While the current implementation is fully functional, potential future enhancements could include support for multiple day forecasts, additional weather metrics like wind speed and humidity, configuration file support for default cities, and more sophisticated output formatting options.

This project serves as an excellent example of how Go's simplicity and powerful standard library can be leveraged to create practical, real-world applications that provide genuine utility to users.
