terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.5.0"
    }
    planetscale = {
      source  = "koslib/planetscale"
      version = "~> 0.5"
    }
  }
}

# # planetscale ##################################################################

# variable "PLANETSCALE_SERVICE_TOKEN_ID" {} # See env.example.
# variable "PLANETSCALE_SERVICE_TOKEN" {}    # See env.example.

provider "planetscale" {
  service_token_id = var.PLANETSCALE_SERVICE_TOKEN_ID
  service_token    = var.PLANETSCALE_SERVICE_TOKEN
}

resource "planetscale_database" "db" {
  organization = "emivespa"
  name         = "db"
}

# output "planetscale_database_html_url" {
#   value = planetscale_database.db.html_url
# }

# aws ##########################################################################

provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

provider "aws" {
  region = "us-east-1"
}

# ecr repo
resource "aws_ecrpublic_repository" "mutant" {
  provider        = aws.us_east_1
  repository_name = "mutant"
}
# output "mutant_ecr_uri" {
#   value = aws_ecrpublic_repository.mutant.repository_uri
# }

# ECR cluster config mostly from here:
# <https://earthly.dev/blog/deploy-dockcontainers-to-awsecs-using-terraform/>
# -- some modifications.

# cluster
resource "aws_ecs_cluster" "mutant" {
  name = "mutant"
}

# variable "DATABASE_URL" {} # See env.example.

# task
resource "aws_ecs_task_definition" "app_task" {
  family                   = "mutant-task"
  container_definitions    = <<DEFINITION
  [
    {
      "name": "mutant-task",
      "image": "${aws_ecrpublic_repository.mutant.repository_uri}:latest",
      "essential": true,
      "environment": [
	{
	  "name": "DATABASE_URL",
	  "value": "${var.DATABASE_URL}"
	}
      ],
      "portMappings": [
        {
          "containerPort": 8080,
          "hostPort": 8080
        }
      ],
      "memory": 512,
      "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = aws_iam_role.myEcsTaskExecutionRole.arn
}

# policy
resource "aws_iam_role" "myEcsTaskExecutionRole" {
  name               = "myEcsTaskExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}
data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}
resource "aws_iam_role_policy_attachment" "ecsTaskExecutionRole_policy" {
  role       = aws_iam_role.myEcsTaskExecutionRole.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# vpc
resource "aws_default_vpc" "default_vpc" {
}
resource "aws_default_subnet" "default_subnet_a" {
  availability_zone = "us-east-1a"
}
resource "aws_default_subnet" "default_subnet_b" {
  availability_zone = "us-east-1b"
}

# alb
resource "aws_alb" "application_load_balancer" {
  name               = "mutant-alb"
  load_balancer_type = "application"
  subnets = [
    "${aws_default_subnet.default_subnet_a.id}",
    "${aws_default_subnet.default_subnet_b.id}"
  ]
  security_groups = ["${aws_security_group.load_balancer_security_group.id}"]
}

# alb sg
resource "aws_security_group" "load_balancer_security_group" {
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lb_target_group" "target_group" {
  name        = "target-group"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_default_vpc.default_vpc.id
}
resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_alb.application_load_balancer.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}

# service
resource "aws_ecs_service" "app_service" {
  name            = "app-first-service"
  cluster         = aws_ecs_cluster.mutant.id
  task_definition = aws_ecs_task_definition.app_task.arn
  launch_type     = "FARGATE"
  desired_count   = 1
  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn # Reference the target group
    container_name   = aws_ecs_task_definition.app_task.family
    container_port   = 8080
  }
  network_configuration {
    subnets          = ["${aws_default_subnet.default_subnet_a.id}", "${aws_default_subnet.default_subnet_b.id}"]
    assign_public_ip = true # Do provide the containers with public IPs
    security_groups  = ["${aws_security_group.service_security_group.id}"]
  }
}

# service sg
resource "aws_security_group" "service_security_group" {
  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = ["${aws_security_group.load_balancer_security_group.id}"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# output "app_url" {
#   value = aws_alb.application_load_balancer.dns_name
# }
