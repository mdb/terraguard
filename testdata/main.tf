data "template_file" "greeting_one" {
  template = <<-EOT
  #/bin/bash
  echo "${var.greeting_one}"
  EOT
}

resource "local_file" "greeting_one" {
  content  = data.template_file.greeting_one.rendered
  filename = "${path.module}/greeting_one.sh"
}

data "template_file" "greeting_two" {
  template = <<-EOT
  #/bin/bash
  echo "${var.greeting_two}"
  EOT
}

resource "local_file" "greeting_two" {
  content  = data.template_file.greeting_two.rendered
  filename = "${path.module}/greeting_two.sh"
}
