# Platform Pilot

<img alt="MetricSample.png" src="https://github.com/NhatMNTran/platform-pilot/blob/master/MetricSample.png?raw=true" data-hpc="true" style="max-width:100%">

AI agent friendly internal developer platform built with **Golang, TypeScript, Kubernetes, and Crossplane**.

---

**Overview**

PlatformPilot is a demo platform engineering system designed to showcase:

* Cloud-native backend development (Golang)
* Developer tooling (TypeScript CLI)
* Kubernetes orchestration
* Infrastructure as Code using Crossplane
* AI agent friendly tooling design (structured contracts)

The system is intentionally designed to behave like a simplified internal developer platform (IDP).

---

**Key Concept**

PlatformPilot **should** allows both humans and AI agents to:

* Create environments
* Deploy services
* Query system status
* Provision infrastructure (via Crossplane)
* Retrieve structured observability data

---

**Architecture**

```text
CLI (TypeScript)
      ↓
Backend API (Golang)
      ↓
Kubernetes Cluster
      ↓
Crossplane Controllers
      ↓
Cloud Providers (AWS/GCP/Azure)
```

---

**Project Structure**

```text
platform-pilot/
│
├── backend/          # Golang API service
├── cli/              # TypeScript CLI tool
├── k8s/              # Kubernetes manifests
├── crossplane/       # Infrastructure definitions
├── docs/             # Architecture + ADRs + contracts
├── .github/          # CI/CD pipelines
└── README.md
```

---

# ⚙️ Tech Stack
* Backend: Golang (Gin)         
* CLI: TypeScript (Node.js) 
* Orchestration: Kubernetes           
* IaC: Crossplane           
* Observability: Prometheus           
* CI/CD: GitHub Actions       

---

**Core Features**

1. Environment Management

Create and manage environments via API or CLI.

```bash
platformctl env-status --id env-123
```

---

2. Service Deployment

Deploy applications into Kubernetes clusters.

```bash
platformctl deploy service --name payments-api
```

---

3. AI-Agent Friendly Tooling

All CLI and API outputs support structured JSON:

```json
{
  "success": true,
  "environment": {
    "id": "env-123",
    "status": "READY"
  }
}
```

---

4. Infrastructure as Code (Crossplane)

Provision cloud resources declaratively:

* S3 Buckets
* Databases
* Networking resources

---

5. Observability

Expose Prometheus metrics:

```text
GET /metrics
```

---

**Running the Project**

1. Backend

```bash
cd backend
go run main.go
```

---

2. CLI

```bash
cd cli
npm install
npm run start -- env-status --id env-123
```

---

3. Kubernetes

```bash
kubectl apply -f k8s/
```

---

4. Crossplane

```bash
kubectl apply -f crossplane/
```

---

**Design Principles**

1. Deterministic Outputs

All system outputs are structured JSON to ensure:

* Automation reliability
* AI agent compatibility
* Predictable system behavior

---

2. Platform Thinking

The system is designed as a reusable platform, not a single application.

---

3. Infrastructure as Code First

All infrastructure is declarative via Kubernetes + Crossplane.

---

4. AI Native Tooling

Every interface is designed to be consumed by:

* Humans
* Scripts
* AI agents

---

**Overview** 

This project demonstrates:

* Cloud-native architecture understanding
* Kubernetes operational knowledge
* Infrastructure automation design
* AI-agent aware system design
* Distributed system thinking

---

**Example Workflow**

```text
1. CLI request → create environment
2. API validates request
3. Kubernetes schedules deployment
4. Crossplane provisions infrastructure
5. Metrics exposed via /metrics
6. AI agent consumes structured output
```

---

**Final Notes**

This project is intentionally designed to be:

* simple enough to build in stages
* advanced enough to demonstrate senior engineering ability
* structured for portfolio and interview storytelling
