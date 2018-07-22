variable "default_site_aws_access_key" {}
variable "default_site_aws_secret_key" {}

variable "site" {
  type = "map"
  default = {
    fqdn = "n.2p5.xyz"
    bucket = "n.2p5.xyz"
    name = "n.2p5.xyz"
    environment = "prod"
    project = "n.2p5.xyz"
    root_object = "index.html"
  }
}