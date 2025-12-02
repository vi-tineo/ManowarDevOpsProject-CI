# Manowar DevOps Project  

Este projeto demonstra um fluxo completo de **CI/CD com GitOps** para aplicações containerizadas em **Kubernetes**.  

## Repositórios  
- **Aplicação e pipeline de build/testes:** [ManowarImage-CI](https://github.com/vi-tineo/ManowarImage-CI)  
- **Repositório GitOps monitorado pelo Argo CD:** [ManowarK8sGitOps-CI](https://github.com/vi-tineo/ManowarK8sGitOps-CI)  

## Fluxo  
1. CI compila o código-fonte, executa testes unitários e gera a imagem Docker  
2. A imagem é publicada no Docker Hub com tag baseada no commit  
3. O workflow atualiza o repositório GitOps com a nova versão  
4. O Argo CD detecta a mudança e aplica no cluster bare metal  

**Representação do fluxo:**  
`[App Repo] → [CI Pipeline] → [Docker Hub] ↓ [GitOps Repo] → [Argo CD] → [Cluster]`  

## Tecnologias  
- GitHub Actions  
- Docker & DockerHub  
- Kubernetes  
- Argo CD  
- GitOps  
- Go  

## Nota Importante  
Este portfólio apresenta um fluxo **simplificado** de CI/CD e GitOps para fins de demonstração.  

Em ambientes de produção, seriam incorporados recursos adicionais como:  
- **Segurança segregada** (RBAC, políticas de rede, gestão de segredos)  
- **Observabilidade** (monitoramento, métricas, tracing e logging estruturado)  
- **Service Mesh** (controle de tráfego, segurança e visibilidade entre serviços)  
- **Infraestrutura como Código (IaC)** com Terraform e providers como AWS/Helm  

Esses componentes garantem confiabilidade, segurança e visibilidade operacional em escala.  
