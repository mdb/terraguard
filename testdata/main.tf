data "template_file" "greeting" {
  template = <<-EOT
  #/bin/bash
  echo "${var.greeting}"
  EOT
}

resource "local_file" "greeting_one" {
  content  = data.template_file.greeting.rendered
  filename = "${path.module}/greeting_one.sh"
}

resource "local_file" "greeting_two" {
  content  = data.template_file.greeting.rendered
  filename = "${path.module}/greeting_two.sh"
}

data "template_file" "greeting_no_changes" {
  template = <<-EOT
  #/bin/bash
  echo "no_changes"
  EOT
}

resource "local_file" "greeting_no_changes" {
  content  = data.template_file.greeting_no_changes.rendered
  filename = "${path.module}/greeting_no_changes.sh"
}
