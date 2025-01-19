# Relicense

Relicense is a simplified licensing platform designed to streamline and manage the licensing process. With a highly customizable and easy-to-scale architecture, Relicense empowers you to efficiently manage licenses according to your unique requirements.

## Technical Features

### Event-Driven Architecture

Relicense leverages an event-driven architecture using RabbitMQ. This enables seamless communication between services, ensuring optimal performance even during peak usage periods. Events are utilized for various operations such as create new requests, updates, and notifications.

### Microservice Architecture

The platform is built upon a microservices architecture to isolate development experiences and make each service individually scalable. This modular approach enhances flexibility and allows for easier management of complex systems.

### Caching with Redis

Redis is employed for caching data and storing user sessions. By caching frequently accessed data, Relicense reduces latency and improves overall system responsiveness. Additionally, Redis facilitates session management, enhancing user experience and scalability.

### High-Performance HTTP Server

Relicense utilizes a high-performance HTTP server powered by the Go HTTP standard library and lightweight chi router. This combination ensures fast and efficient handling of HTTP requests, optimizing system performance and response times.

### Containerized Development and Deployment

Development and deployment processes are simplified with Docker containerization. By encapsulating each component into containers, Relicense enables consistent and portable deployments across different environments. This approach enhances scalability, reliability, and ease of management.

### Comprehensive Testing

To ensure reliability and quality, Relicense employs comprehensive automated testing methodologies. Unit tests and integration tests are conducted throughout the development lifecycle, validating the functionality and performance of the application.

### OAuth Authentication

Relicense supports OAuth2 authentication for seamless authorization processes. This feature allows users to securely log in using their existing OAuth provider credentials. OAuth authentication enhances security and user experience by eliminating the need for separate login credentials and simplifying access control.

## Business Features

### Create Relicense Template

Administrators have the ability to create customizable templates for license requests. These templates can be tailored to specific license types or organizational needs, allowing for streamlined and standardized request processes.

### Apply For A License

Users can submit requests for new licenses through the platform. These requests are then routed to administrators for review and approval. Users can track the status of their requests and receive updates on progress.

### License Approval

Administrators have the responsibility to review and approve license requests. They can assess whether the request meets the necessary criteria and either approve it for issuance or reject it with feedback if it does not comply with requirements.

### Monitoring Request

Administrators have access to a dashboard where they can monitor all active and submitted license requests. This feature provides visibility into the status of each request, allowing administrators to prioritize and manage them efficiently.

### Email Notification

The platform automatically sends email notifications to users upon approval or rejection of their license requests. This feature ensures timely communication and keeps users informed about the status of their requests without requiring them to constantly check the platform.

## Getting Started

To get started with Relicense, follow these steps:

1. Install Docker and Docker Compose on your development environment.
2. Clone the Relicense repository from GitHub.
3. Configure environment variables and settings according to your deployment environment.
4. Build Docker images for each microservice component.
5. Run "docker compose up" to deploy or "docker compose watch" to develop the Relicense platform.
6. Access the platform via the provided HTTP endpoints or user interface.
7. Customize and scale the platform as needed to meet your licensing requirements.

## Contributions

Relicense welcomes contributions from the community. If you would like to contribute to the development of the platform, please refer to the contribution guidelines in the repository.

## License

Relicense is released under the MIT License. See the `LICENSE` file for more information.
