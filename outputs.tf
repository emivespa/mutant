output "planetscale_database_html_url" {
  value = planetscale_database.db.html_url
}

output "mutant_ecr_uri" {
  value = aws_ecrpublic_repository.mutant.repository_uri
}

output "app_url" {
  value = aws_alb.application_load_balancer.dns_name
}
