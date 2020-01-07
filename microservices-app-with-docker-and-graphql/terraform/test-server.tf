module "test-server" {
  source = "./node-server"

  ami-id   = "ami-00068cd7555f543d5"
  key-pair = aws_key_pair.microservices-demo-key.key_name
  name     = "Test server"
}
