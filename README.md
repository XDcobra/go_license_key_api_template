# GoFiber Starter Stack

🚀 A production-ready Go Fiber microservices template with Docker Compose orchestration, featuring Redis clustering, MySQL database, Prometheus monitoring, and Grafana visualization.

---

## 🚀 Features

- **Go Fiber API Gateway** - High-performance HTTP framework with built-in middleware
- **Redis Cluster** - Master-slave replication with Redis Sentinel for high availability
- **MySQL Database** - Relational database with GORM ORM integration
- **Prometheus** - Metrics collection and monitoring
- **Grafana** - Data visualization and dashboards
- **Redis Insight** - Redis GUI for database management
- **Docker Compose** - Complete container orchestration
- **Makefile** - Simplified development commands

---

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.24.4 or higher
- Git

---

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Gateway   │    │     Grafana     │    │   Prometheus    │
│   (Go Fiber)    │    │   (Port 3000)   │    │   (Port 9090)   │
│   (Port 8000)   │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
         ┌─────────────────────────────────────────────────┐
         │              Redis Cluster                      │
         │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ │
         │  │  Master │ │  Slave  │ │Sentinel1│ │Sentinel2│ │
         │  │ (6379)  │ │ (6379)  │ │ (26379) │ │ (26380) │ │
         │  └─────────┘ └─────────┘ └─────────┘ └─────────┘ │
         │                    │                    │         │
         │              ┌─────────┐         ┌─────────┐     │
         │              │Sentinel3│         │Redis    │     │
         │              │(26381)  │         │Insight  │     │
         │              └─────────┘         │(5540)   │     │
         └──────────────────────────────────┴─────────┴─────┘
                                 │
         ┌─────────────────────────────────────────────────┐
         │              MySQL Database                     │
         │              (Port 3306)                        │
         └─────────────────────────────────────────────────┘
```

---

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/XDcobra/gofiber-starter-stack.git
   cd gofiber-starter-stack
   ```

2. **Initialize Go module (if needed)**
   ```bash
   cd services/api_gateway
   go mod init github.com/XDcobra/gofiber-starter-stack
   go mod tidy
   cd ../..
   ```

3. **Create environment file**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Start all services**
   ```bash
   make docker-start
   ```

5. **Access the services**
   - API Gateway: http://localhost:8000
   - Grafana: http://localhost:3000
   - Prometheus: http://localhost:9090
   - Redis Insight: http://localhost:5540
   - MySQL: localhost:3306


6. **For Production: Review Security settings**
   - Please read section [🔐 Authentication & Security](#-authentication--security) and change the passwords (and usernames)
within the .env file, as well as reviewing the exposed service ports

---

## 📁 Project Structure

```
gofiber-starter-stack/
├── docker-compose.yml              # Main orchestration file
├── Makefile                        # Development commands
├── README.md                       # This file
├── .env.example                    # Environment variables template
└── services/
    ├── api_gateway/                # Go Fiber API Gateway
    │   ├── controller/             # HTTP controllers
    │   │   ├── DummyController/    # Example controller
    │   │   ├── MySQLController/    # MySQL operations
    │   │   └── RedisController/    # Redis operations
    │   ├── database/               # Database connections
    │   │   ├── MySQL/              # MySQL connection & models
    │   │   └── Redis/              # Redis connection
    │   ├── model/                  # Data models
    │   ├── prometheus/             # Metrics configuration
    │   ├── router/                 # Route definitions
    │   ├── services/               # Business logic
    │   ├── Dockerfile              # API Gateway container
    │   ├── go.mod                  # Go dependencies
    │   └── main.go                 # Application entry point
    ├── grafana/                    # Grafana configuration
    ├── mysql/                      # MySQL initialization
    ├── prometheus/                 # Prometheus configuration
    └── dummy_service*/             # Example microservices that you can add in the future
```

---

## 🔧 API Endpoints

### Health Check
- `GET /` - API health check

### Redis Operations
- `GET /redis/ping` - Redis connection test
- `GET /redis/get` - Get value from Redis
- `POST /redis/post` - Set value in Redis

### MySQL Operations
- `GET /mysql/get/:id` - Get record by ID
- `POST /mysql/post` - Create new record

### Prometheus Metrics
- `GET /metrics` - Get the raw prometheus exported data

---

## 🔐 Authentication & Security

### Password Protected Endpoints

| Service | Endpoint               | Authentication Type | Environment Variable                                       | Description                     |
|---------|------------------------|---------------------|------------------------------------------------------------|---------------------------------|
| **Prometheus** | http://localhost:9090  | Basic Auth          | `PROM_USER` / `PROM_PASS`                                  | Metrics collection dashboard    |
| **Grafana** | http://localhost:3000  | Admin Login         | `GF_SECURITY_ADMIN_USER` / `GF_SECURITY_ADMIN_PASSWORD` | Data visualization dashboard    |
| **Redis Insight** | http://localhost:5540  | None                | -                                                          | Redis GUI (no auth by default)  |
| **API Gateway** | http://localhost:8000/metrics | Basic Auth          | `METRICS_USER` / `METRICS_PASS`                            | Exported metrics for prometheus |
| **MySQL** | localhost:3306         | Database Auth       | `MYSQL_ROOT_PASSWORD` / `MYSQL_USER` / `MYSQL_PASSWORD`    | Database connection             |

### Environment Variables for Authentication

```env
# Golang API Gateway metrics endpoint
METRICS_USER='metrics_user'
METRICS_PASS='metrics_password'

# Prometheus Authentication
PROM_USER='prometheus_user'
PROM_PASS='prometheus_password'

# Grafana Authentication  
GF_SECURITY_ADMIN_USER='admin'
GF_SECURITY_ADMIN_PASSWORD='password'

# MySQL Authentication
MYSQL_ROOT_PASSWORD='root_password'
MYSQL_USER='user'
MYSQL_PASSWORD='password'

# MySQL Database name
MYSQL_DATABASE='example_db'
```

### Default Credentials

- **Prometheus**: Use the values from `PROM_USER` and `PROM_PASS` in your `.env` file
- **Grafana**: 
  - Username: `admin`
  - Password: Value of `GF_SECURITY_ADMIN_PASSWORD` in your `.env` file
- **MySQL**: 
  - Root Password: Value of `MYSQL_ROOT_PASSWORD` in your `.env` file
  - User Password: Value of `MYSQL_PASSWORD` in your `.env` file

### Security Notes

- **API Gateway endpoints** are currently public - implement authentication middleware for production
- **Redis Insight** has no authentication by default - consider adding reverse proxy with auth
- **Prometheus and Grafana** use basic authentication - ensure strong passwords in production
- **MySQL** uses database-level authentication - keep credentials secure
- **Docker** exposes all service ports at the moment - for production builds, consider to only expose the service ports
that should be reachable from the outside (e.g. remove redis-sentinel from being reachable from outside as it is not needed
normally)

---

## 🐳 Docker Services

| Service | Port | Description |
|---------|------|-------------|
| API Gateway | 8000 | Go Fiber application |
| Grafana | 3000 | Data visualization |
| Prometheus | 9090 | Metrics collection |
| Redis Master | 6379 | Redis primary instance |
| Redis Slave | 6379 | Redis replica |
| Redis Sentinel 1 | 26379 | Redis sentinel for HA |
| Redis Sentinel 2 | 26380 | Redis sentinel for HA |
| Redis Sentinel 3 | 26381 | Redis sentinel for HA |
| Redis Insight | 5540 | Redis GUI |
| MySQL | 3306 | Database |

---

## 🛠️ Development Commands

```bash
# Start all services
make docker-start

# Stop all services
make docker-stop

# Stop and remove containers
make docker-down

# View logs
docker-compose logs -f

# Rebuild specific service
docker-compose build api_gateway
```

---

## 📊 Monitoring & Observability

### Prometheus Metrics
The API Gateway automatically exposes metrics at `/metrics` endpoint:
- HTTP request duration
- Request count by status code
- Active connections
- Custom business metrics

In case you want to expose more custom metrics, consider to add these into the [Gofiber-Prometheus](./services/api_gateway/prometheus/fiberprometheus.go) file

### Grafana Dashboards
Pre-configured dashboards for:
- API Gateway performance
- Redis cluster health
- MySQL database metrics
- System resource usage

---

## 🔄 Redis Cluster Configuration

The Redis setup includes:
- **Master-Slave Replication** for data redundancy
- **Redis Sentinel** for automatic failover
- **Redis Insight** for visual management
- **High Availability** with automatic master election

---

## 🗄️ Database Setup

### MySQL
- Automatic schema migration with GORM
- Connection pooling
- Transaction support
- Model auto-generation

### Redis
- Connection pooling
- Automatic reconnection
- Sentinel support for HA
- Pub/Sub capabilities

---

## 🚀 Production Deployment

1. **Update environment variables** for production
2. **Configure SSL/TLS** certificates
3. **Set up proper logging** and log rotation
4. **Configure backup strategies** for databases
5. **Set up monitoring alerts** in Grafana
6. **Use external volumes** for data persistence

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/) - Fast HTTP framework
- [GORM](https://gorm.io/) - Go ORM library
- [Redis](https://redis.io/) - In-memory data store
- [Prometheus](https://prometheus.io/) - Monitoring system
- [Grafana](https://grafana.com/) - Data visualization

---

## 📞 Support

If you have any questions or need help, please open an issue on GitHub.

---

**Happy Coding! 🎉**
