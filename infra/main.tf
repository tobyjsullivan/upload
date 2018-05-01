terraform {
  backend "s3" {
    bucket = "terraform-states.tobyjsullivan.com"
    key = "states/upload/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = "ap-southeast-2"
}

resource "aws_security_group" "allow_udp_in" {
  ingress {
    from_port = 80
    protocol = "udp"
    to_port = 80
  }
}

resource "aws_key_pair" "login" {
  key_name_prefix = "rapidupload"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCob5Cr3R33z+jiVQVyhtBNMKR5TpAjnfRDSANckdcoiE0U4gN/aWBatbGF9vlM1ywIZRdimbKC0+/08kFiYd4vJQzB4EHwMG0LK2UF85dqPLhonZeeNVe1XmEQAvSrzzOi33CbbiBIP3AWNyw85xvBIoVRRdRb8ryf10DRmroX0rVEgWTdTxQj44bPn1yD6MIjzFZzbUFUgdM7hnvBz4UhucCcv9vEqA+QLT+DWyGtAOU03ZjRrvRn+lGqbrv/lZ2IbQ8W1nNQ3A/utZT8nuShNk/s0gBikrsanF9RrXY+bpBSas4hRV2ukmMC89Hj+65J4jV8Kj75lEEkjOXmcDI2Nbv5eiHLUbcK40q9pfnFx1Ufvmn+0H8jQY5VF2p7N03a3+knmVsUpcwC80CC81iSIqYuQDxyZsF9hZD6Rd4BJBhSsw1GJ1akDcZmpJ33PO89+GSEtRmi68w/0sgTYlEbAEOYLe+mNthv/RI24GrRHF/6ZAVKcY+BT61z+yR60p2dRHhos8NAJt+WV/jk4YvEsa1PisfJ1Bnkc29jkNK4qjQv+uuXEhrLQurOItamxuQYvSx++rZhvNwYMqpDgCgPzwvHYnUo6BVFyoiPVymJbY3PkOR7kQFu5xUMPiKzDN1Tdg54VPwdoh2ymD3DWOJ2le5m776Wf9bkbfJbRMG6AQ== tobyjsullivan@gmail.com"
}

resource "aws_instance" "server" {
  ami = "ami-d38a4ab1"
  instance_type = "c5.large"

  security_groups = ["${aws_security_group.allow_udp_in.name}"]
  key_name = "${aws_key_pair.login.key_name}"
}

output "ip_address" {
  value = "${aws_instance.server.public_ip}"
}
