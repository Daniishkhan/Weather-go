
# GoWeather

GoWeather is a powerful yet simple weather application built with Go. It provides current weather information and forecasts through both a RESTful API and a CLI interface.

## Features

- Fetch current weather data for any location
- Get weather forecasts for up to 7 days
- User authentication for personalized experiences
- Store and manage user preferences (default location, temperature unit)
- Cache weather data to minimize API calls
- Track user search history
- CLI interface for quick weather checks
- Dockerized for easy deployment


## Technology Stack

- **Backend**: Go 1.22+ with Chi framework
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Configuration**: Viper
- **Logging**: Zap
- **CLI**: Cobra
- **HTTP Client**: Resty
- **Caching**: Go-cache (in-memory caching)
- **API Documentation**: Swaggo
- **Containerization**: Docker

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    default_location VARCHAR(100),
    preferred_unit VARCHAR(1) CHECK (preferred_unit IN ('C', 'F')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE search_history (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    location VARCHAR(100) NOT NULL,
    searched_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## API Endpoints

- `POST /api/register` - Register a new user
- `POST /api/login` - Login and receive JWT
- `GET /api/weather?location={city}` - Get current weather
- `GET /api/forecast?location={city}&days={number}` - Get forecast
- `GET /api/preferences` - Get user preferences
- `PUT /api/preferences` - Update user preferences
- `GET /api/history` - Get user's search history

## CLI Commands

- `goweather login <username> <password>`
- `goweather weather <location>`
- `goweather forecast <location> [days]`
- `goweather set-preferences [--location <default_location>] [--unit <C|F>]`
- `goweather history`

## Project Structure

```
goweather/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers.go
│   ├── auth/
│   │   └── auth.go
│   ├── weather/
│   │   └── service.go
│   └── user/
│       └── service.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── postgres.go
│   ├── cache/
│   │   └── cache.go
│   └── weatherapi/
│       └── client.go
├── docker-compose.yml
└── Dockerfile
```

This README provides an overview of the project, its features, technology stack, database schema, API endpoints, CLI commands, and basic setup instructions. It gives potential users and contributors a clear understanding of what the project does and how it's structured.
