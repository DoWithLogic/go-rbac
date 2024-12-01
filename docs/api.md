## About

The **go-rbac** repository is a robust, Golang-based system for implementing and managing role and scope based access control (RBAC) within your organization. It provides a framework to define user roles, permissions, and associated access controls to protect resources and ensure that only authorized users can access specific features or perform certain actions.

## Overview

- **Purpose:** The go-rbac system is designed to simplify the management of user roles, permissions, and access controls within your application, making it easier to implement security policies.
- **Technology Stack:** Developed in Golang, leveraging lightweight libraries to ensure high performance and scalability.
- **Data Management:** Implements user roles, permissions, and role-permission mappings in a structured way to easily enforce security policies.
- **Scalability:** Supports multiple roles and permissions, designed to scale with growing organizational needs.
- **Security:** Ensures fine-grained access control with role-based permissions and JWT authentication, safeguarding critical resources.
- **Integration:** Easily integrates with existing authentication systems, APIs, and services within your organization.

## Key Components

- **User Role Model:** Defines the structure and attributes of user roles, including standard roles like Admin, Employee, and Customer.
- **Permissions Model:** Defines the various permissions (scopes) users can have based on their roles.
- **Role-Permission Mapping:** Links specific roles to allowed permissions, ensuring users with the appropriate roles can perform specific actions.
- **JWT Authentication:** Secure access to the API using JWT tokens, with support for role and permission validation.
- **API Endpoints:** Exposes various endpoints to manage users, roles, and permissions, ensuring access control via role validation.

## API Documentation

### Overview
This repository provides a set of API endpoints for managing roles, permissions, and user access. The API allows you to create, update, retrieve, and delete roles, permissions, and role-permission mappings. It also supports secure JWT-based authentication to enforce role-based access control.

### Explore Swagger Documentation
For a detailed description of all the available API endpoints, request/response formats, and examples, explore our Swagger documentation at the following link:

- [Swagger Documentation](http://localhost:3002/swagger/index.html)

The Swagger documentation will provide detailed information on:
- **Available Endpoints**: All API routes for managing users, roles, permissions, and access control.
- **Request/Response Formats**: Detailed format for the expected input and output of each API endpoint.
- **Authentication**: How to authenticate requests using JWT tokens.
- **Role and Permission Validation**: How roles and permissions are validated for each endpoint.

### Contact

For any questions or support related to go-rbac, please create issue contact at `martin.yonathan305@gmail.com`.

### Swagger Documentation

Explore our Swagger documentation for a comprehensive overview of the available endpoints, request/response formats, and examples. Access the documentation at `https://{{base-url}}/swagger/index.html`.

