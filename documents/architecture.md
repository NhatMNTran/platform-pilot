# PlatformPilot Architecture

## Overview
Utilized technologies:
- Golang backend API
- TypeScript CLI tooling
- Kubernetes orchestration
- Crossplane infrastructure management

## System Design

CLI → Go API → Kubernetes → Crossplane → Cloud Providers

## Key Principles
- Structured JSON contracts for automation
- Kubernetes-native infrastructure management
- AI-agent consumable tooling interfaces
- Declarative infrastructure over imperative scripts

## Data Flow
1. CLI sends structured command
2. Go API validates request
3. Kubernetes executes deployment
4. Crossplane provisions infrastructure
5. Status returned as JSON