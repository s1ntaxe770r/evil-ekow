module "app_deploy" {
  source  = "kbst.xyz/catalog/custom-manifests/kustomization"
  version = "0.1.0"

  configuration_base_key = "prod"
  configuration = {
    prod = {
      namespace = "app-${terraform.workspace}"
      resources = [
        "${path.root}/manifests/namespace.yaml",
        "${path.root}/manifests/deployment.yaml",
        "${path.root}/manifests/service.yaml",
      ]

      common_labels = {
        "env" = terraform.workspace
      }
    }
    stage = {}
  }
}
