# GoFiber Starter Stack

🚀 A production-ready Go Fiber microservices template with multiple deployment options, 
featuring Redis clustering, MySQL database, comprehensive monitoring, and observability solutions.

---

## 🎯 What This Repository Offers

This repository provides a complete microservices stack built with **Go Fiber** that can be deployed using either:

- **🐳 Docker Compose** - For local development and simple deployments
- **☸️ Kubernetes with Helm** - For production-ready cloud deployments

### 🌟 Key Features

- **Go Fiber API Gateway** - High-performance HTTP framework with built-in middleware
- **Redis Cluster** - Master-slave replication with Redis Sentinel for high availability
- **MySQL Database** - Relational database with GORM ORM integration
- **Comprehensive Monitoring Stack**:
  - **Prometheus** - Metrics collection and monitoring
  - **Grafana** - Data visualization and dashboards
  - **Loki & Promtail** - Log aggregation and querying (Kubernetes only)
- **Redis Insight** - Redis GUI for database management
- **Go Swagger Documentation** - Auto-generated API documentation
- **Prometheus Auto Service Discovery** - Automatic service monitoring (Kubernetes only)

---

## 🚀 Quick Start

Choose your deployment method:

### 🐳 Docker Compose (Local Development)
Perfect for local development, testing, and simple deployments.

📖 **[Click here for Docker Compose Documentation](./services/README.md)**

```bash
# Quick start with Docker Compose
git clone https://github.com/XDcobra/gofiber-starter-stack.git
cd gofiber-starter-stack
make docker-start
```

### ☸️ Kubernetes with Helm (Production)
Enterprise-ready deployment with advanced features like log aggregation and auto service discovery.

📖 **[Click here for Kubernetes/Helm Documentation](./charts/README.md)**

```bash
# Quick start with Kubernetes
git clone https://github.com/XDcobra/gofiber-starter-stack.git
cd gofiber-starter-stack/charts/gofiber-starter-stack
helm install gofiber-starter-stack .
```

---

## 🏗️ Architecture Overview

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

         ┌─────────────────────────────────────────────────┐
         │              Loki & Promtail                   │
         │              (Kubernetes Only)                 │
         │  ┌─────────┐ ┌─────────┐ ┌─────────┐           │
         │  │  Loki   │ │Promtail1│ │Promtail2│           │
         │  │ (3100)  │ │         │ │         │           │
         │  └─────────┘ └─────────┘ └─────────┘           │
         └─────────────────────────────────────────────────┘
```

---

## 📋 Prerequisites

### For Docker Compose
- Docker and Docker Compose
- Go 1.24.4 or higher
- Git

### For Kubernetes/Helm
- Kubernetes cluster (minikube, kind, or cloud provider)
- Helm 3.x
- kubectl configured
- Ingress controller (nginx-ingress recommended)

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

### Documentation & Monitoring
- `GET /swagger/*` - Swagger API documentation
- `GET /metrics` - Prometheus metrics endpoint

---

## 🔐 Security Features

Both deployment methods include:

- **Authentication** for monitoring dashboards (Prometheus, Grafana)
- **Database security** with configurable credentials
- **Service isolation** and network security
- **Environment-based configuration** management

---

## 📊 Monitoring & Observability

### Docker Compose
- Prometheus metrics collection
- Grafana dashboards
- Basic logging

### Kubernetes/Helm
- **Enhanced monitoring** with auto service discovery
- **Centralized logging** with Loki & Promtail
- **Advanced dashboards** and alerting
- **Service mesh ready** architecture

---

## 🚀 Production Deployment

### Docker Compose
- Suitable for small to medium deployments
- Easy to set up and maintain
- Good for development and testing environments

### Kubernetes/Helm
- **Enterprise-grade** scalability and reliability
- **Auto-scaling** and load balancing
- **Advanced monitoring** and logging
- **Multi-environment** support (dev, staging, prod)

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
- [Loki](https://grafana.com/oss/loki/) - Log aggregation
- [Helm](https://helm.sh/) - Kubernetes package 
- [Go Swagger](https://github.com/gofiber/swagger) - Swagger for GoFiber

---

## 📞 Support

If you have any questions or need help, please open an issue on GitHub.
