resource "aws_instance" "default" {
  ami                  = var.ami-id
  iam_instance_profile = var.iam-instance-profile
  instance_type        = var.instance-type
  # name                   = var.name
  key_name = var.key-pair
  # key_pair_key           = var.key-pair-key
  private_ip             = var.private-ip
  subnet_ip              = var.subnet-ip
  vpc_security_group_ids = var.vpc-security-group-ids

  tags = {
    Name = var.name
  }
}
