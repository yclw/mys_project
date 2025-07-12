# 基本配置
SHELL := /bin/bash
.DEFAULT_GOAL := help

# 项目信息
PROJECT_NAME := mys_project
API_SERVICE := apps/api
USER_SERVICE := apps/user
BIN_DIR := bin

# Go相关配置
GO := go
GOBUILD := go build
GOWORK := go work
GOMOD := go mod

# Docker相关
DOCKER_COMPOSE := docker-compose
DOCKER_COMPOSE_FILE := deployments/docker-compose.yml

# 颜色定义
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

help: ## 显示帮助信息
	@echo -e "$(BLUE)$(PROJECT_NAME) Makefile$(RESET)"
	@echo -e "$(BLUE)==================$(RESET)"
	@echo ""
	@echo -e "$(GREEN)可用命令:$(RESET)"
	@echo -e "  $(YELLOW)make init$(RESET)    - 初始化项目环境"
	@echo -e "  $(YELLOW)make pb$(RESET)      - 生成protobuf代码"
	@echo -e "  $(YELLOW)make run$(RESET)     - 运行项目"

init: ## 初始化项目环境
	@echo -e "$(GREEN)正在初始化项目环境...$(RESET)"
	@mkdir -p $(BIN_DIR) logs
	@mkdir -p apps/api/logs apps/user/logs
	@$(GOWORK) sync
	@echo -e "$(GREEN)0. 初始化项目...$(RESET)"
	@echo -e "$(GREEN)1. 初始化API服务...$(RESET)"
	@cd $(API_SERVICE) && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)2. 初始化用户服务...$(RESET)"
	@cd $(USER_SERVICE) && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)3. 初始化公共包...$(RESET)"
	@cd pkg/common && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)4. 初始化protobuf包...$(RESET)"
	@cd pkg/protobuf && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)5. 初始化utils包...$(RESET)"
	@cd pkg/utils && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)6. 初始化model包...$(RESET)"
	@cd pkg/model && $(GOMOD) download && $(GOMOD) tidy
	@echo -e "$(GREEN)项目环境初始化完成$(RESET)"

pb: ## 生成protobuf代码
	@echo -e "$(GREEN)正在生成protobuf代码...$(RESET)"
	@cd pkg/protobuf && chmod +x scripts/generate.sh && ./scripts/generate.sh
	@echo -e "$(GREEN)protobuf代码生成完成$(RESET)"

run: ## 运行项目
	@echo -e "$(GREEN)正在启动项目...$(RESET)"
	@make kill
	@echo -e "$(GREEN)1. 启动基础设施...$(RESET)"
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d mysql redis etcd
	@echo -e "$(GREEN)2. 构建服务...$(RESET)"
	@cd $(USER_SERVICE) && $(GOBUILD) -o $(shell pwd)/$(BIN_DIR)/user ./main.go
	@cd $(API_SERVICE) && $(GOBUILD) -o $(shell pwd)/$(BIN_DIR)/api ./main.go
	@echo -e "$(GREEN)3. 启动用户服务...$(RESET)"
	@cd $(USER_SERVICE) && $(shell pwd)/$(BIN_DIR)/user &
	@sleep 3
	@echo -e "$(GREEN)4. 启动API服务...$(RESET)"
	@cd $(API_SERVICE) && $(shell pwd)/$(BIN_DIR)/api &
	@echo -e "$(GREEN)项目启动完成！$(RESET)"
	@echo -e "$(BLUE)API服务: http://localhost:8080$(RESET)"
	@echo -e "$(BLUE)用户服务: grpc://localhost:8081$(RESET)"

kill: ## 杀死进程
	@echo -e "$(GREEN)正在杀死进程...$(RESET)"
	@pkill -f $(BIN_DIR)/api || true
	@pkill -f $(BIN_DIR)/user || true
	@echo -e "$(GREEN)进程杀死完成$(RESET)"

.PHONY: help init pb run kill